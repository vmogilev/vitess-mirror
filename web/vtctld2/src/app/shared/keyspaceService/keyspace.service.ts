import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions } from '@angular/http';
import { Keyspace } from '../keyspaceObject/keyspace/'
import {Observable} from 'rxjs/Observable';


@Injectable()
export class KeyspaceService {
  private keyspacesUrl = '../api/keyspaces/';
  private srv_keyspaceUrl = '../api/srv_keyspace/';
  private keyspace_shardsUrl = '../api/shards/';
  vtTabletTypes = [
    'unknown', 'master', 'replica', 'rdonly', 'spare', 'experimental',
    'backup', 'restore', 'worker'
  ];
  constructor(private http: Http) {}

  getShards(keyspaceName) {
    return this.http.get(this.keyspace_shardsUrl + keyspaceName + "/")
    .map( (resp) => {
      return resp.json(); 
    })
  }
  getKeyspaceNames() {
    return this.http.get(this.keyspacesUrl)
    .map( (resp) => {
      return resp.json();
    });
  }
  getSrvKeyspaces() {
    return this.http.get(this.srv_keyspaceUrl + "local/")
    .map( (resp) => {
      return resp.json();
    });
  }
  getSeperatedShards(keyspaceName, partition) {
    return this.getShards(keyspaceName)
      .map(allShards =>{
        let keyspace = {};
        let shardSet = {};
        keyspace["name"] = keyspaceName;
        keyspace["servingShards"] = [];
        keyspace["nonservingShards"] = [];
        let shardReferences = partition.shard_references;
        if (shardReferences != undefined) {
          shardReferences.forEach( servingShard => {
            keyspace["servingShards"].push(servingShard.name);
            shardSet[servingShard.name] = true;
          });
        }
        allShards.forEach( shard => {
          if (!(shard in shardSet)) {
            keyspace["nonservingShards"].push(shard);
          }
        });
        return keyspace;
      })
      .map( (keyspace) => {
        return this.getKeyspaceRaw(keyspace["name"])
          .map(keyspaceData => {
            keyspace["shardingColumnName"] = "sharding_column_name" in keyspaceData ? keyspaceData.sharding_column_name: "";
            keyspace["shardingColumnType"] = "sharding_column_type" in keyspaceData ? keyspaceData.sharding_column_type : "";
            return keyspace;
          });
      });
  }
  SrvKeyspaceAndNamesObservable(){
    let keyspaceNamesStream = this.getKeyspaceNames();
    let srvKeyspaceStream = this.getSrvKeyspaces();
    return keyspaceNamesStream.combineLatest(srvKeyspaceStream);
  }
  getKeyspaces() {
    return this.SrvKeyspaceAndNamesObservable()
    .map( (streams) => {
        let keyspaceNames = streams[0];
        if(keyspaceNames.length < 1) return [];
        let srvKeyspace = streams[1];
        let allDone = null;
        keyspaceNames.forEach( keyspaceName => {
          let partitions = [];
          if (srvKeyspace[keyspaceName] == undefined) {
            partitions = [{served_type: 1}];
          } else {
            partitions = srvKeyspace[keyspaceName].partitions;
          }
          for (let p = 0; p < partitions.length; p++) {
            let partition = partitions[p];
            if (this.vtTabletTypes[partition.served_type] == 'master') {
              let shardStream = this.getSeperatedShards(keyspaceName, partition);
              if (allDone == null) {
                allDone = shardStream;
              } else {
                allDone = allDone.merge(shardStream);
              }
              break;
            }
          }
        });
        return allDone;
      }
    );
  }
  getKeyspace(keyspaceName) {
    return this.getSrvKeyspaces()
    .map( (srvKeyspace) => {
        let allDone = null;
        let partitions = [];
        if (srvKeyspace[keyspaceName] == undefined) {
          partitions = [{served_type: 1}];
        } else {
          partitions = srvKeyspace[keyspaceName].partitions;
        }
        for (let p = 0; p < partitions.length; p++) {
          let partition = partitions[p];
          if (this.vtTabletTypes[partition.served_type] == 'master') {
            return this.getSeperatedShards(keyspaceName, partition);
          }
        }
      }
    );
  }
  getKeyspaceRaw(keyspaceName) {
    return this.http.get(this.keyspacesUrl + keyspaceName)
    .map( (resp) => {
      return resp.json();
    });
  }
  sendPostRequest(url, body) {
    let headers = new Headers({ 'Content-Type': 'application/x-www-form-urlencoded' });
    let options = new RequestOptions({ headers: headers });
    return this.http.post(url, body, options)
    .map( (resp) => {
      return resp.json(); 
    });
  }
  createKeyspace(keyspace) {
    let body = "action=CreateKeyspace" + "&shardingColumnName=" + keyspace.shardingColumnName + "&shardingColumnType=" + keyspace.shardingColumnType;
    return this.sendPostRequest(this.keyspacesUrl + keyspace.name, body);
  }
  deleteKeyspace(keyspace) {
    let body = "action=DeleteKeyspace";
    return this.sendPostRequest(this.keyspacesUrl + keyspace.name, body);
  }
  editKeyspace(keyspace) {
    let body = "action=EditKeyspace" + "&shardingColumnName=" + keyspace.shardingColumnName + "&shardingColumnType=" + keyspace.shardingColumnType;
    return this.sendPostRequest(this.keyspacesUrl + keyspace.name, body);
  }
}
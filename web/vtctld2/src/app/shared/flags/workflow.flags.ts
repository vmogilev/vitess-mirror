import { CheckBoxFlag, DropDownFlag, InputFlag } from './flag';

export class NewWorkflowFlags {
  flags= {};
  constructor(workflows) {
    // General flags.
    this.flags['skip_start'] = new SkipStartFlag(0, 'skip_start');
    this.flags['skip_start'].positional = true;
    this.flags['skip_start'].namedPositional = 'skip_start';
    this.flags['factory_name'] = new FactoryNameFlag(1, 'factory_name', workflows);
    this.flags['factory_name'].positional = true;

    // Flags for the Sleep workflow.
    this.flags['sleep_duration'] = new SleepDurationFlag(2, 'sleep_duration');
    this.flags['sleep_duration'].positional = true;
    this.flags['sleep_duration'].namedPositional = 'duration';

    // Flags for the Schema Swap workflow.
    this.flags['schema_swap_keyspace'] = new SchemaSwapKeyspaceFlag(3, 'schema_swap_keyspace');
    this.flags['schema_swap_keyspace'].positional = true;
    this.flags['schema_swap_keyspace'].namedPositional = 'keyspace';
    this.flags['schema_swap_sql'] = new SchemaSwapSQLFlag(4, 'schema_swap_sql');
    this.flags['schema_swap_sql'].positional = true;
    this.flags['schema_swap_sql'].namedPositional = 'sql';

    // Flags for Horizontal Resharding workflow
    this.flags['horizontal_resharding_keyspace'] = new HorizontalReshardKeyspaceFlag(5, 'horizontal_resharding_keyspace');
    this.flags['horizontal_resharding_keyspace'].positional = true;
    this.flags['horizontal_resharding_keyspace'].namedPositional = 'keyspace';
    this.flags['horizontal_resharding_sources'] = new HorizontalReshardSourceFlag(6, 'horizontal_resharding_sources');
    this.flags['horizontal_resharding_sources'].positional = true;
    this.flags['horizontal_resharding_sources'].namedPositional = 'source_shard_list';
    this.flags['horizontal_resharding_destinations'] = new HorizontalReshardDestinationFlag(7, 'horizontal_resharding_destinations');
    this.flags['horizontal_resharding_destinations'].positional = true;
    this.flags['horizontal_resharding_destinations'].namedPositional = 'destination_shard_list';
    this.flags['horizontal_resharding_vtworker'] = new HorizontalReshardVTWorkerFlag(8, 'horizontal_resharding_vtworker');
    this.flags['horizontal_resharding_vtworker'].positional = true
    this.flags['horizontal_resharding_vtworker'].namedPositional = 'vtworker_server_address'
  }
}

export class FactoryNameFlag extends DropDownFlag {
  constructor(position: number, id: string, workflows) {
    super(position, id, 'Factory Name', 'Specifies the type of workflow to create.', '');
    let options = [];
    // tested the UI; TO DO: load workflow.horizontal_resharding component
    if (workflows.horizontal_resharding) {
      options.push({
        label: 'Horizontal Resharding',
        value: 'horizontal_resharding'
      });
    }   
    if (workflows.schema_swap) {
      options.push({
        label: 'Schema Swap',
        value: 'schema_swap'
      });
    }
    if (workflows.sleep) {
      options.push({
        label: 'Sleep',
        value: 'sleep'
      });
    }
    if (workflows.topo_validator) {
      options.push({
        label: 'Topology Validator',
        value: 'topo_validator',
      });
    }
    this.setOptions(options);
    this.value = options[0].value;
  }
}

export class SkipStartFlag extends CheckBoxFlag {
  constructor(position: number, id: string, value= false) {
    super(position, id, 'Skip Start', 'Create the workflow, but don\'t start it.', value);
  }
}

export class SleepDurationFlag extends InputFlag {
  constructor(position: number, id: string, value= '') {
    super(position, id, 'Sleep Duration', 'Time to sleep for, in seconds.', value);
    this.setDisplayOn('factory_name', 'sleep');
  }
}

export class SchemaSwapKeyspaceFlag extends InputFlag {
  constructor(position: number, id: string, value= '') {
    super(position, id, 'Keyspace', 'Name of a keyspace.', value);
    this.setDisplayOn('factory_name', 'schema_swap');
  }
}

export class SchemaSwapSQLFlag extends InputFlag {
  constructor(position: number, id: string, value= '') {
    super(position, id, 'SQL', 'SQL representing the schema change.', value);
    this.setDisplayOn('factory_name', 'schema_swap');
  }
}

export class HorizontalReshardKeyspaceFlag extends InputFlag {
  constructor(position: number, id: string, value= '') {
    super(position, id, 'Keyspace', 'Name of a keyspace.', value);
    this.setDisplayOn('factory_name', 'horizontal_resharding');
  }
}

export class HorizontalReshardDestinationFlag extends InputFlag {
  constructor(position: number, id: string, value= '') {
    super(position, id, 'Destination Shards', 'List of destination shards.', value);
    this.setDisplayOn('factory_name', 'horizontal_resharding');
  }
} 

export class HorizontalReshardSourceFlag extends InputFlag {
  constructor(position: number, id: string, value= '') {
    super(position, id, 'Source Shards', 'List of source shards.', value);
    this.setDisplayOn('factory_name', 'horizontal_resharding');
  }
} 

export class HorizontalReshardVTWorkerFlag extends InputFlag {
  constructor(position: number, id: string, value= '') {
    super(position, id, 'VTWorker Server', 'server address of vtworker.', value);
    this.setDisplayOn('factory_name', 'horizontal_resharding');
  }
} 

// WorkflowFlags is used by the Start / Stop / Delete dialogs.
export class WorkflowFlags {
  flags= {};
  constructor(path) {
    this.flags['workflow_uuid'] = new WorkflowUuidFlag(0, 'workflow_uuid', path);
    this.flags['workflow_uuid']['positional'] = true;
  }
}

export class WorkflowUuidFlag extends InputFlag {
  constructor(position: number, id: string, value= '') {
    super(position, id, '', '', value, false);
  }
}

set @@sql_mode='';
drop table if exists onlineddl_test;
create table onlineddl_test (
  id int auto_increment,
  i int not null,
  dt datetime not null default '1970-00-00 00:00:00',
  primary key(id)
) auto_increment=1;

drop event if exists onlineddl_test;
delimiter ;;
create event onlineddl_test
  on schedule every 1 second
  starts current_timestamp
  ends current_timestamp + interval 60 second
  on completion not preserve
  enable
  do
begin
  insert into onlineddl_test values (null, 11, now());
  insert into onlineddl_test values (null, 13, now());
  insert into onlineddl_test values (null, 17, now());
end ;;

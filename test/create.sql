create table amz.test (
     did    integer primary key default nextval('serial'),
     name   varchar(40) not null check (name <> '')
);

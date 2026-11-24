-- Create a new database
create database hoc_golang;

-- Drop a database
drop database hoc_golang;

-- Create a new schema
create schema school;

-- Drop a schema
drop schema school cascade;

-- One - One
-- Create user table
create table if not exists users (
	user_id serial primary key,
	name varchar(50) not null,
	email  varchar(50) not null 
);

-- Create profile table
create table if not exists profiles (
	profile_id serial primary key,
	user_id int unique not null,
	phone varchar(15),
	address varchar(100),
	constraint fk_user foreign key (user_id) references users(user_id) on delete cascade
);

-- Drop table (dangerours)
drop table if exists profiles;
drop table if exists users;


-- One - Many
-- Create Categories table
create table if not exists categories (
	category_id serial primary key,
	name varchar(50) not null
);

-- Create products table
create table if not exists products (
	product_id serial primary key,
	category_id int not null,
	name varchar(100) not null,
	price int not null check (price > 0),
	image varchar(255) not null,
	status int not null check (status in (1,2)),
	constraint fk_category foreign key (category_id) references categories (category_id) on delete restrict
);

-- Drop table (dangerours)
drop table if exists products;
drop table if exists categories;


-- Many - Many
-- Create Students table
create table if not exists students (
	student_id serial primary key,
	name varchar(50) not null
);

-- Create Courses table
create table if not exists courses (
	course_id serial primary key,
	course_name varchar(50) not null
);

-- Create Students-Course table
create table if not exists students_courses (
	student_id int not null,
	course_id int not null,
	primary key(student_id, course_id),
	constraint fk_student foreign key (student_id) references students (student_id) on delete cascade,
	constraint fk_course foreign key (course_id) references courses (course_id) on delete cascade
);

-- Drop table (dangerours)
drop table if exists students_courses;
drop table if exists courses;
drop table if exists students;


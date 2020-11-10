-- public.userprofile definition

-- Drop table

-- DROP TABLE public.userprofile;

CREATE TABLE public.userprofile (
	email varchar(100) NOT NULL,
	CONSTRAINT userprofile_pkey PRIMARY KEY (email)
);


-- public.block definition

-- Drop table

-- DROP TABLE public.block;

CREATE TABLE public.block (
	id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	requestor varchar(100) NULL,
	target varchar(100) NULL,
	CONSTRAINT block_pkey PRIMARY KEY (id),
	CONSTRAINT block_requestor_fkey FOREIGN KEY (requestor) REFERENCES userprofile(email),
	CONSTRAINT block_target_fkey FOREIGN KEY (target) REFERENCES userprofile(email)
);


-- public.friend definition

-- Drop table

-- DROP TABLE public.friend;

CREATE TABLE public.friend (
	id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	emailuserone varchar(100) NULL,
	emailusertwo varchar(100) NULL,
	CONSTRAINT friend_pkey PRIMARY KEY (id),
	CONSTRAINT friend_emailuserone_fkey FOREIGN KEY (emailuserone) REFERENCES userprofile(email),
	CONSTRAINT friend_emailusertwo_fkey FOREIGN KEY (emailusertwo) REFERENCES userprofile(email)
);


-- public."subscription" definition

-- Drop table

-- DROP TABLE public."subscription";

CREATE TABLE public."subscription" (
	id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	requestor varchar(100) NULL,
	target varchar(100) NULL,
	CONSTRAINT subscription_pkey PRIMARY KEY (id),
	CONSTRAINT subscription_requestor_fkey FOREIGN KEY (requestor) REFERENCES userprofile(email),
	CONSTRAINT subscription_target_fkey FOREIGN KEY (target) REFERENCES userprofile(email)
);
-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id serial4 NOT NULL,
	"name" varchar(255) NOT NULL,
	phone_number varchar(20) NOT NULL,
	nationality_id varchar(20) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);


-- public.financial_accounts definition

-- Drop table

-- DROP TABLE public.financial_accounts;

CREATE TABLE public.financial_accounts (
	id int4 DEFAULT nextval('user_transactions_id_seq'::regclass) NOT NULL,
	user_id int4 NOT NULL,
	balance numeric(15, 2) DEFAULT 0.00 NOT NULL,
	bank_account_number varchar(50) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	CONSTRAINT user_transactions_pkey PRIMARY KEY (id),
	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "branches" (
    "branch_id" UUID PRIMARY KEY,
    "branch_code" VARCHAR(50) NOT NULL,
    "branch_name" VARCHAR(50) NOT NULL,
    "branch_address" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "curiers" (
    "curier_id" UUID PRIMARY KEY,
    "first_name" VARCHAR(50) NOT NULL,
    "last_name" VARCHAR(50) NOT NULL,
    "phone_number" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
)

CREATE TABLE IF NOT EXISTS "employees" (
    "employee_id" UUID PRIMARY KEY,
    "store_id" UUID PRIMARY KEY,
    "first_name" VARCHAR(50) NOT NULL,
    "last_name" VARCHAR(50) NOT NULL,
    "phone_number" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
)

CREATE TABLE IF NOT EXISTS "stores" (
    "store_id" UUID PRIMARY KEY,
    "branch_id" UUID PRIMARY KEY,
    "store_name" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
)

SELECT
	COUNT(*) OVER(),
	branch_id,
	branch_code,
	branch_name,
	branch_address,
	TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
	TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS'),
	COALESCE(TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS'), '')
FROM "branches"
WHERE COALESCE(TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS'), '')=''
AND TRUE ORDER BY created_at DESC OFFSET 0 

SELECT
	COUNT(*) OVER(),
	curier_id,
	first_name,
	last_name,
	phone_number,
	TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
	TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS'),
	COALESCE(TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS'), '')
FROM "curiers"
WHERE COALESCE(TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS'), '')=''
AND TRUE ORDER BY created_at DESC OFFSET 0 

SELECT
	COUNT(*) OVER(),
	employee_id,
    store_id
	first_name,
	last_name,
	phone_number,
	TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
	TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS'),
	COALESCE(TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS'), '')
FROM "branches"
WHERE COALESCE(TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS'), '')=''
AND TRUE ORDER BY created_at DESC OFFSET 0 

SELECT
	COUNT(*) OVER(),
	store_id,
	branch_id,
	store_name,
	TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
	TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS'),
	COALESCE(TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS'), '')
FROM "branches"
WHERE COALESCE(TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS'), '')=''
AND TRUE ORDER BY created_at DESC OFFSET 0 
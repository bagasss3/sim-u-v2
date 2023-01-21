-- +goose Up

CREATE TYPE status_type AS ENUM (
  'PENDING',
  'ACTIVE',
  'NONACTIVE'
);

CREATE TYPE user_gender AS ENUM (
  'MALE',
  'FEMALE'
);

CREATE TYPE user_role AS ENUM (
  'ADMIN',
  'STUDENT_ORGANIZATION',
  'STUDENT'
);

CREATE TABLE IF NOT EXISTS users (
  id BIGINT PRIMARY KEY,
  name TEXT,
  email text,
  password text,
  role user_role,
  created_at timestamp NOT NULL DEFAULT 'now()',
  updated_at timestamp NOT NULL DEFAULT 'now()',
  deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS sessions (
    id BIGINT PRIMARY KEY,
    access_token TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    refresh_token_expired_at timestamp NOT NULL,
    user_id INTEGER NOT NULL,
    created_at timestamp NOT NULL DEFAULT 'now()',
    updated_at timestamp NOT NULL DEFAULT 'now()',
    deleted_at timestamp
);

ALTER TABLE sessions ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE UNIQUE INDEX IF NOT EXISTS email_unique_index ON users(email);

CREATE TABLE IF NOT EXISTS images (
    id BIGINT NOT NULL PRIMARY KEY,
    public_id VARCHAR(191) NOT NULL,
    width INTEGER NOT NULL,
    height INTEGER NOT NULL,
    version INTEGER NOT NULL,
    format VARCHAR(191) NOT NULL,
    etag VARCHAR(191) NOT NULL,
    url VARCHAR(191) NOT NULL,
    secure_url VARCHAR(191) NOT NULL,
    signature VARCHAR(191) NOT NULL,
    created_at timestamp NOT NULL DEFAULT 'now()',
    updated_at timestamp NOT NULL DEFAULT 'now()',
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS student_organizations (
    id BIGINT NOT NULL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    image_id BIGINT NULL,
    description TEXT NULL,
    is_verified status_type,
    created_at timestamp NOT NULL DEFAULT 'now()',
    updated_at timestamp NOT NULL DEFAULT 'now()',
    deleted_at timestamp
);

ALTER TABLE student_organizations ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE student_organizations ADD FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE SET NULL ON UPDATE CASCADE;

CREATE UNIQUE INDEX IF NOT EXISTS student_organization_user_id_key ON student_organizations(user_id);

CREATE TABLE IF NOT EXISTS student_organization_preconditions (
    id BIGINT NOT NULL PRIMARY KEY,
    student_organization_id BIGINT NOT NULL,
    image_id BIGINT NOT NULL,
    created_at timestamp NOT NULL DEFAULT 'now()',
    updated_at timestamp NOT NULL DEFAULT 'now()',
    deleted_at timestamp
);

ALTER TABLE student_organization_preconditions ADD FOREIGN KEY (student_organization_id) REFERENCES student_organizations(id) ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE student_organization_preconditions ADD FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE SET NULL ON UPDATE CASCADE;

CREATE UNIQUE INDEX IF NOT EXISTS Student_Organization_Precondition_student_organization_id_key ON student_organization_preconditions(student_organization_id);

CREATE UNIQUE INDEX IF NOT EXISTS Student_Organization_Precondition_image_id_key ON student_organization_preconditions(image_id);

CREATE TABLE IF NOT EXISTS students (
    id BIGINT NOT NULL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    gender user_gender,
    birth_date timestamp NULL,
    image_id BIGINT NULL,
    status status_type,
    created_at timestamp NOT NULL DEFAULT 'now()',
    updated_at timestamp NOT NULL DEFAULT 'now()',
    deleted_at timestamp
);

ALTER TABLE students ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE students ADD FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE SET NULL ON UPDATE CASCADE;

CREATE UNIQUE INDEX IF NOT EXISTS Student_user_id_key ON students(user_id);

CREATE UNIQUE INDEX IF NOT EXISTS Student_phone_number_key ON students(phone_number);

CREATE TABLE categories (
    id BIGINT NOT NULL PRIMARY KEY,
    name VARCHAR(191) NOT NULL,
    description TEXT,
    created_at timestamp NOT NULL DEFAULT 'now()',
    created_by BIGINT NOT NULL,
    updated_at timestamp NOT NULL DEFAULT 'now()',
    updated_by BIGINT NOT NULL,
    deleted_at timestamp,
    deleted_by BIGINT
);

-- CreateTable
CREATE TABLE eligibilities (
    id BIGINT NOT NULL PRIMARY KEY,
    name VARCHAR(191) NOT NULL,
    description TEXT,
    created_at timestamp NOT NULL DEFAULT 'now()',
    created_by BIGINT NOT NULL,
    updated_at timestamp NOT NULL DEFAULT 'now()',
    updated_by BIGINT NOT NULL,
    deleted_at timestamp,
    deleted_by BIGINT
);

CREATE UNIQUE INDEX IF NOT EXISTS Category_name_key ON categories(name);

CREATE UNIQUE INDEX IF NOT EXISTS Eligibility_name_key ON eligibilities(name);

CREATE TABLE IF NOT EXISTS events (
    id BIGINT NOT NULL PRIMARY KEY,
    student_organization_id BIGINT NOT NULL,
    name VARCHAR(191) NOT NULL,
    date_time timestamp NOT NULL,
    image_id BIGINT NULL,
    description TEXT NOT NULL,
    precondition INTEGER NOT NULL DEFAULT 0,
    is_verified status_type,
    category_id BIGINT NOT NULL,
    eligibility_id BIGINT NOT NULL,
    created_at timestamp NOT NULL DEFAULT 'now()',
    created_by BIGINT NOT NULL,
    updated_at timestamp NOT NULL DEFAULT 'now()',
    updated_by BIGINT NOT NULL,
    deleted_at timestamp,
    deleted_by BIGINT
);

ALTER TABLE events ADD FOREIGN KEY (student_organization_id) REFERENCES student_organizations(id) ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE events ADD FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE SET NULL ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS eventstudents (
    id BIGINT NOT NULL PRIMARY KEY,
    student_id BIGINT NOT NULL,
    event_id BIGINT NOT NULL,
    created_at timestamp NOT NULL DEFAULT 'now()',
    updated_at timestamp NOT NULL DEFAULT 'now()',
    deleted_at timestamp
);

ALTER TABLE eventstudents ADD FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE eventstudents ADD FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE SET NULL ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS eventpreconditions (
    id BIGINT NOT NULL PRIMARY KEY,
    event_student_id BIGINT NOT NULL,
    image_id BIGINT NOT NULL,
    created_at timestamp NOT NULL DEFAULT 'now()',
    created_by BIGINT NOT NULL,
    updated_at timestamp NOT NULL DEFAULT 'now()',
    updated_by BIGINT NOT NULL,
    deleted_at timestamp,
    deleted_by BIGINT
);

ALTER TABLE eventpreconditions ADD FOREIGN KEY (event_student_id) REFERENCES eventstudents(id) ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE eventpreconditions ADD FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE SET NULL ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS eventcomments (
    id BIGINT NOT NULL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    event_id BIGINT NOT NULL,
    comment VARCHAR(255) NOT NULL,
    parent_id BIGINT NOT NULL,
    mentioned_user_ids BIGINT[] DEFAULT '{}'::BIGINT[],
    created_at timestamp NOT NULL DEFAULT 'now()',
    created_by BIGINT NOT NULL,
    updated_at timestamp NOT NULL DEFAULT 'now()',
    updated_by BIGINT NOT NULL,
    deleted_at timestamp,
    deleted_by BIGINT
);

ALTER TABLE eventcomments ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE eventcomments ADD FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE RESTRICT ON UPDATE CASCADE;

-- +goose Down
DROP TABLE IF EXISTS eventcomments;
DROP TABLE IF EXISTS eventpreconditions;
DROP TABLE IF EXISTS eventstudents;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS eligibilities;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS student_organization_preconditions;
DROP TABLE IF EXISTS student_organizations;
DROP TABLE IF EXISTS images;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS status_type;
DROP TYPE IF EXISTS user_gender;
DROP TYPE IF EXISTS user_role;
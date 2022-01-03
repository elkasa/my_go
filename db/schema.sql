-- idfs
CREATE TABLE idfs (
  idf_id text PRIMARY KEY ,
  Created_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  updated_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  direct                     text NOT NULL,
  version                          text    DEFAULT 'v1.0' NOT NULL,
  Is_Enabled     BOOLEAN DEFAULT TRUE NOT NULL,
  Config         JSONB DEFAULT '{}'::JSONB NOT NULL,
   CHECK (char_length(idf_id)>=8 AND char_length(idf_id)<=10)
);
-- partenaires
CREATE TABLE parts (
  part_id text PRIMARY KEY ,
  Created_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  updated_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  version                          text    DEFAULT 'v1.0' NOT NULL,
  Is_Enabled     BOOLEAN DEFAULT TRUE NOT NULL,
  Config         JSONB DEFAULT '{}'::JSONB NOT NULL,
   CHECK (char_length(part_id)>=8 AND char_length(part_id)<=10)
);

-- profile
CREATE TABLE profiles (
  profile_name text PRIMARY KEY ,
  profile_path text UNIQUE NOT NULL ,
  Created_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  updated_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  version                          text    DEFAULT 'v1.0' NOT NULL,
  environment                      text    DEFAULT 're7' NOT NULL,
  Is_Enabled     BOOLEAN DEFAULT TRUE NOT NULL,
   CHECK (char_length(profile_name)>=8 AND char_length(profile_name)<=30)
);
-- buckets
CREATE TABLE buckets (
  bucket_name text PRIMARY KEY ,
  Created_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  updated_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  profile_name           text    NOT NULL REFERENCES profiles(profile_name),
  idf_id                 text    NOT NULL REFERENCES idfs(idf_id),
  part_id                text    NOT NULL REFERENCES parts(part_id),
  Is_Enabled     BOOLEAN DEFAULT TRUE NOT NULL,
  Config         JSONB DEFAULT '{}'::JSONB NOT NULL,
  version                          text    DEFAULT 'test_bucket' NOT NULL,
CHECK (char_length(bucket_name)>=8 AND char_length(bucket_name)<=50)
);
DO
$$
BEGIN
   IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'tellerdb') THEN
      CREATE DATABASE tellerdb;
   END IF;
END
$$;
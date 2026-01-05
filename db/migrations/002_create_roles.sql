CREATE ROLE housing_app
    LOGIN
    PASSWORD '1111';

GRANT CONNECT ON DATABASE housing_db TO housing_app;
GRANT USAGE ON SCHEMA public TO housing_app;
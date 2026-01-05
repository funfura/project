SELECT 
    current_user,
    current_database();

ROLLBACK;

SELECT code FROM work_types;

SELECT column_name, data_type
FROM information_schema.columns
WHERE table_name = 'work_types';

DROP TABLE IF EXISTS work_types CASCADE;

SELECT system_code, scope, code, name
FROM work_types
ORDER BY system_code, code;

SELECT system_code, scope, code, name
FROM work_types
ORDER BY system_code, scope, code;

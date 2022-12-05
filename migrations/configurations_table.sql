-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS configurations (
    id serial PRIMARY KEY,
	package VARCHAR ( 100 ) NOT NULL,
	country_code VARCHAR ( 2 ) NOT NULL,
	percentile_min int NOT NULL,
	percentile_max int NOT NULL,
    main_sku VARCHAR ( 100 ) NOT NULL
    );

CREATE INDEX IF NOT EXISTS package_country_idx
ON configurations(package,country_code);

INSERT INTO
    configurations (package,country_code,percentile_min,percentile_max,main_sku)
VALUES
    ('com.softinit.iquitos.mainapp', 'US', 0, 25, 'rdm_premium_v3_020_trial_7d_monthly'), 
    ('com.softinit.iquitos.mainapp', 'US', 25, 50, 'rdm_premium_v3_030_trial_7d_monthly'), 
    ('com.softinit.iquitos.mainapp', 'US', 50, 75, 'rdm_premium_v3_100_trial_7d_yearly'), 
    ('com.softinit.iquitos.mainapp', 'US', 75, 100, 'rdm_premium_v3_150_trial_7d_yearly'), 
    ('com.softinit.iquitos.mainapp', 'ZZ', 0, 100, 'rdm_premium_v3_050_trial_7d_yearly');


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE configurations;
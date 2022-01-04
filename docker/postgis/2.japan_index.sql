select 'Add Index to japan'::text;
CREATE INDEX japan_pref_idx     ON japan (pref);
CREATE INDEX japan_regional_idx ON japan (regional);
CREATE INDEX japan_city1_idx    ON japan (city1);
CREATE INDEX japan_city2_idx    ON japan (city2);
CREATE INDEX japan_code_idx     ON japan (code);
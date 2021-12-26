select 'Add Index to japan'::text;
CREATE
INDEX japan_pref_idx     ON japan (pref);
CREATE
INDEX japan_regional_idx ON japan (regional);
CREATE
INDEX japan_city1_idx    ON japan (city1);
CREATE
INDEX japan_city2_idx    ON japan (city2);
CREATE
INDEX japan_code_idx     ON japan (code);

select 'Create prefectures'::text;
DROP TABLE IF EXISTS public.prefectures;
CREATE TABLE public.prefectures
(
    gid  serial PRIMARY KEY,
    code varchar(2)  not null,
    name varchar(10) not null,
    geom geometry(polygon, 4612)
);
CREATE
INDEX prefectures_geom_idx ON public.prefectures USING gist (geom);


do $$ declare p record;
begin
    for p in
        select substring(min(code), 0, 3) as code,
               pref                       as name
        from japan
        group by pref
        order by code
    loop
        raise info 'prefecture % %', p.code, p.name;
        insert into prefectures
        (
         code,
         name,
         geom
        )
        select
            p.code,
            p.name,
            geom
        from (
                 select
                    (ST_Dump(ST_Union(geom))).geom as geom
                 from
                     japan
                 where
                     pref = p.name
                   and
                     ST_Area(geom) > 0.0003495285322129
             ) as tbl;
    end loop;
end$$;

--select 'Drop japan'::text;
--DROP TABLE IF EXISTS public.japan;

select 'prefectures'::text;
select code, name, count(*) from public.prefectures group by code, name order by code;


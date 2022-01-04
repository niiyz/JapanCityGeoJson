select 'Create cities'::text;
DROP TABLE IF EXISTS public.cities;
CREATE TABLE public.cities
(
    gid  serial PRIMARY KEY,
    code varchar(5)  not null,
    name varchar(50) not null,
    geom geometry(polygon, 4612)
);
CREATE INDEX cities_geom_idx ON public.cities USING gist (geom);
CREATE INDEX cities_code_idx ON public.cities (code);


do $$ declare p record;
begin
    for p in
        select
             code,
             concat(pref, regional, city1, city2) as name
        from
             japan
        group by pref, regional, city1, city2, code
        order by code
    loop
        raise info 'city % %', p.code, p.name;
        insert into cities
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
                     code = p.code
                 and
                     ST_Area(geom) > 0.0003495285322129
             ) as tbl;
    end loop;
end$$;

--select 'Drop japan'::text;
--DROP TABLE IF EXISTS public.japan;

select 'cities'::text;
select count(*) from public.cities;


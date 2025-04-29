SELECT id_murid, name, status as pendidikan_terakhir, m.time_create, MAX(p.time_create) as time_updated
FROM pendidikan p
INNER JOIN murid m on p.id_murid = m.id
GROUP BY name
ORDER BY id_murid
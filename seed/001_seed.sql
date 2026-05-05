-- DRIVERS (imágenes de Wikipedia/wikimedia, cargan sin CORS)
INSERT INTO drivers (name, team, nationality, number, image_url) VALUES
('Max Verstappen',    'Oracle Red Bull Racing',       'Dutch',      1,  'https://upload.wikimedia.org/wikipedia/commons/thumb/6/67/Max_Verstappen_2023_%28cropped%29.jpg/440px-Max_Verstappen_2023_%28cropped%29.jpg'),
('Lando Norris',      'McLaren Formula 1 Team',       'British',    4,  'https://upload.wikimedia.org/wikipedia/commons/thumb/f/f7/Lando_Norris_2023_%28cropped%29.jpg/440px-Lando_Norris_2023_%28cropped%29.jpg'),
('Charles Leclerc',   'Scuderia Ferrari',             'Monegasque', 16, 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/91/Charles_Leclerc_2024_%28cropped%29.jpg/440px-Charles_Leclerc_2024_%28cropped%29.jpg'),
('Carlos Sainz',      'Scuderia Ferrari',             'Spanish',    55, 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/72/Carlos_Sainz_Jr._2023_%28cropped%29.jpg/440px-Carlos_Sainz_Jr._2023_%28cropped%29.jpg'),
('Lewis Hamilton',    'Mercedes-AMG Petronas',        'British',    44, 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/18/Lewis_Hamilton_2016_Malaysia_2.jpg/440px-Lewis_Hamilton_2016_Malaysia_2.jpg'),
('George Russell',    'Mercedes-AMG Petronas',        'British',    63, 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a9/George_Russell_%28driver%29_2023_%28cropped%29.jpg/440px-George_Russell_%28driver%29_2023_%28cropped%29.jpg'),
('Fernando Alonso',   'Aston Martin Aramco',          'Spanish',    14, 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e2/Fernando_Alonso_2023_%28cropped%29.jpg/440px-Fernando_Alonso_2023_%28cropped%29.jpg'),
('Oscar Piastri',     'McLaren Formula 1 Team',       'Australian', 81, 'https://upload.wikimedia.org/wikipedia/commons/thumb/d/d9/Oscar_Piastri_2023_%28cropped%29.jpg/440px-Oscar_Piastri_2023_%28cropped%29.jpg'),
('Lance Stroll',      'Aston Martin Aramco',          'Canadian',   18, 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/59/Lance_Stroll_2023_%28cropped%29.jpg/440px-Lance_Stroll_2023_%28cropped%29.jpg'),
('Pierre Gasly',      'Alpine F1 Team',               'French',     10, 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/74/Pierre_Gasly_2023_%28cropped%29.jpg/440px-Pierre_Gasly_2023_%28cropped%29.jpg');

-- RACES (imágenes de circuitos, Wikimedia)
INSERT INTO races (grand_prix, circuit, country, race_date, image_url) VALUES
('Bahrain Grand Prix',        'Bahrain International Circuit',        'Bahrain',     '2025-03-02', 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6b/Bahrain_International_Circuit%2C_2004.jpg/640px-Bahrain_International_Circuit%2C_2004.jpg'),
('Saudi Arabian Grand Prix',  'Jeddah Corniche Circuit',              'Saudi Arabia','2025-03-09', 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3d/Jeddah_Street_Circuit.jpg/640px-Jeddah_Street_Circuit.jpg'),
('Australian Grand Prix',     'Albert Park Circuit',                  'Australia',   '2025-03-23', 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/47/Albert_Park_Circuit_aerial.jpg/640px-Albert_Park_Circuit_aerial.jpg'),
('Japanese Grand Prix',       'Suzuka International Racing Course',   'Japan',       '2025-04-06', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a4/Suzuka_circuit_2005.png/640px-Suzuka_circuit_2005.png'),
('Chinese Grand Prix',        'Shanghai International Circuit',       'China',       '2025-04-20', 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6e/Shanghai_International_Circuit.jpg/640px-Shanghai_International_Circuit.jpg'),
('Miami Grand Prix',          'Miami International Autodrome',        'USA',         '2025-05-04', 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/5f/Miami_International_Autodrome.jpg/640px-Miami_International_Autodrome.jpg'),
('Emilia Romagna Grand Prix', 'Autodromo Enzo e Dino Ferrari',        'Italy',       '2025-05-18', 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/28/Autodromo_Enzo_e_Dino_Ferrari.jpg/640px-Autodromo_Enzo_e_Dino_Ferrari.jpg'),
('Monaco Grand Prix',         'Circuit de Monaco',                    'Monaco',      '2025-05-25', 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Monaco_Formula_1_track_map.svg/640px-Monaco_Formula_1_track_map.svg.png'),
('Canadian Grand Prix',       'Circuit Gilles-Villeneuve',            'Canada',      '2025-06-15', 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/50/Circuit_Gilles_Villeneuve.png/640px-Circuit_Gilles_Villeneuve.png'),
('Spanish Grand Prix',        'Circuit de Barcelona-Catalunya',       'Spain',       '2025-06-29', 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/81/Circuit_de_Catalunya.png/640px-Circuit_de_Catalunya.png');

-- RESULTS para Bahrain GP (race_id = 1)
INSERT INTO results (driver_id, race_id, position, points, fastest_lap) VALUES
(1, 1, 1,  25, false),
(2, 1, 2,  18, false),
(3, 1, 3,  15, true),
(4, 1, 4,  12, false),
(5, 1, 5,  10, false),
(6, 1, 6,  8,  false),
(7, 1, 7,  6,  false),
(8, 1, 8,  4,  false),
(9, 1, 9,  2,  false),
(10,1, 10, 1,  false);

-- RESULTS para Saudi GP (race_id = 2)
INSERT INTO results (driver_id, race_id, position, points, fastest_lap) VALUES
(3, 2, 1,  25, false),
(1, 2, 2,  18, true),
(2, 2, 3,  15, false),
(8, 2, 4,  12, false),
(5, 2, 5,  10, false),
(4, 2, 6,  8,  false),
(6, 2, 7,  6,  false),
(7, 2, 8,  4,  false),
(10,2, 9,  2,  false),
(9, 2, 10, 1,  false);

-- RESULTS para Australian GP (race_id = 3)
INSERT INTO results (driver_id, race_id, position, points, fastest_lap) VALUES
(2, 3, 1,  25, true),
(8, 3, 2,  18, false),
(1, 3, 3,  15, false),
(3, 3, 4,  12, false),
(7, 3, 5,  10, false),
(4, 3, 6,  8,  false),
(5, 3, 7,  6,  false),
(6, 3, 8,  4,  false),
(10,3, 9,  2,  false),
(9, 3, 10, 1,  false);
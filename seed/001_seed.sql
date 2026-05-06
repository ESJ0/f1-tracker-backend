TRUNCATE results, drivers, races RESTART IDENTITY CASCADE;

-- DRIVERS
INSERT INTO drivers (name, team, nationality, number, image_url) VALUES
('Max Verstappen',  'Oracle Red Bull Racing', 'Dutch',       1,  'https://ESJ0.github.io/f1-tracker-frontend/assets/drivers/max-verstappen.png'),
('Lando Norris',    'McLaren Formula 1 Team', 'British',     4,  'https://ESJ0.github.io/f1-tracker-frontend/assets/drivers/lando-norris.png'),
('Charles Leclerc', 'Scuderia Ferrari',       'Monegasque',  16, 'https://ESJ0.github.io/f1-tracker-frontend/assets/drivers/charles-leclerc.png'),
('Carlos Sainz',    'Williams Racing',        'Spanish',     55, 'https://ESJ0.github.io/f1-tracker-frontend/assets/drivers/carlos-sainz.png'),
('Lewis Hamilton',  'Scuderia Ferrari',       'British',     44, 'https://ESJ0.github.io/f1-tracker-frontend/assets/drivers/lewis-hamilton.png'),
('George Russell',  'Mercedes-AMG Petronas',  'British',     63, 'https://ESJ0.github.io/f1-tracker-frontend/assets/drivers/george-russell.png'),
('Fernando Alonso', 'Aston Martin Aramco',    'Spanish',     14, 'https://ESJ0.github.io/f1-tracker-frontend/assets/drivers/fernando-alonso.png'),
('Oscar Piastri',   'McLaren Formula 1 Team', 'Australian',  81, 'https://ESJ0.github.io/f1-tracker-frontend/assets/drivers/oscar-piastri.png'),
('Lance Stroll',    'Aston Martin Aramco',    'Canadian',    18, 'https://ESJ0.github.io/f1-tracker-frontend/assets/drivers/lance-stroll.png'),
('Pierre Gasly',    'Alpine F1 Team',         'French',      10, 'https://ESJ0.github.io/f1-tracker-frontend/assets/drivers/pierre-gasly.png');

-- RACES
INSERT INTO races (grand_prix, circuit, country, race_date, image_url) VALUES
('Gran Premio de Bahréin',        'Bahrain International Circuit',      'Bahrain',      '2025-03-02', 'https://ESJ0.github.io/f1-tracker-frontend/assets/races/bahrain.png'),
('Gran Premio de Arabia Saudita', 'Jeddah Corniche Circuit',            'Saudi Arabia', '2025-03-09', 'https://ESJ0.github.io/f1-tracker-frontend/assets/races/saudi-arabia.png'),
('Gran Premio de Australia',      'Albert Park Circuit',                'Australia',    '2025-03-23', 'https://ESJ0.github.io/f1-tracker-frontend/assets/races/australia.png'),
('Gran Premio de Japón',          'Suzuka International Racing Course', 'Japan',        '2025-04-06', 'https://ESJ0.github.io/f1-tracker-frontend/assets/races/japon.png'),
('Gran Premio de China',          'Shanghai International Circuit',     'China',        '2025-04-20', 'https://ESJ0.github.io/f1-tracker-frontend/assets/races/china.png'),
('Gran Premio de Miami',          'Miami International Autodrome',      'USA',          '2025-05-04', 'https://ESJ0.github.io/f1-tracker-frontend/assets/races/miami.png'),
('Gran Premio de Emilia Romagna', 'Autodromo Enzo e Dino Ferrari',      'Italy',        '2025-05-18', 'https://ESJ0.github.io/f1-tracker-frontend/assets/races/italy.png'),
('Gran Premio de Mónaco',         'Circuit de Monaco',                  'Monaco',       '2025-05-25', 'https://ESJ0.github.io/f1-tracker-frontend/assets/races/monaco.png'),
('Gran Premio de Canadá',         'Circuit Gilles-Villeneuve',          'Canada',       '2025-06-15', 'https://ESJ0.github.io/f1-tracker-frontend/assets/races/canada.png'),
('Gran Premio de España',         'Circuit de Barcelona-Catalunya',     'Spain',        '2025-06-29', 'https://ESJ0.github.io/f1-tracker-frontend/assets/races/spain.png');

-- RESULTS Bahréin
INSERT INTO results (driver_id, race_id, position, points, fastest_lap) VALUES
(1,1,1,25,false),(2,1,2,18,false),(3,1,3,15,true),
(4,1,4,12,false),(5,1,5,10,false),(6,1,6,8,false),
(7,1,7,6,false),(8,1,8,4,false),(9,1,9,2,false),(10,1,10,1,false);

-- RESULTS Arabia Saudita
INSERT INTO results (driver_id, race_id, position, points, fastest_lap) VALUES
(3,2,1,25,false),(1,2,2,18,true),(2,2,3,15,false),
(8,2,4,12,false),(5,2,5,10,false),(4,2,6,8,false),
(6,2,7,6,false),(7,2,8,4,false),(10,2,9,2,false),(9,2,10,1,false);

-- RESULTS Australia
INSERT INTO results (driver_id, race_id, position, points, fastest_lap) VALUES
(2,3,1,25,true),(8,3,2,18,false),(1,3,3,15,false),
(3,3,4,12,false),(7,3,5,10,false),(4,3,6,8,false),
(5,3,7,6,false),(6,3,8,4,false),(10,3,9,2,false),(9,3,10,1,false);
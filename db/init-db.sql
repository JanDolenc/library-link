-- create users table
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL
);

-- insert 100 users (data from Mockaroo)
INSERT INTO users (first_name, last_name) VALUES
('Hedwiga', 'Cumpsty'),
('Nikola', 'Wonter'),
('Marleen', 'Brownsell'),
('Everard', 'Casely'),
('Selestina', 'Meaker'),
('Aldous', 'Latour'),
('Fanechka', 'Paur'),
('Silvia', 'Nutton'),
('Adele', 'Macey'),
('Melvyn', 'Giamuzzo'),
('Zorine', 'Saintsbury'),
('Glory', 'Blaksley'),
('Fabiano', 'Lonnon'),
('Lilas', 'Rivitt'),
('Lalo', 'Beagan'),
('Keane', 'Dunk'),
('Cristiano', 'Meeland'),
('Derwin', 'Andren'),
('June', 'Winsome'),
('Ciro', 'Archbutt'),
('Florence', 'Jope'),
('Milo', 'Bertelet'),
('Dinny', 'Bufton'),
('Therese', 'Tidcombe'),
('Sherlock', 'Norwich'),
('Arlyne', 'Strettell'),
('Lin', 'Jickles'),
('Tressa', 'Seden'),
('Bone', 'Perschke'),
('Hirsch', 'Phythean'),
('Rinaldo', 'Dallison'),
('Karim', 'Robbings'),
('Kamillah', 'Trenholm'),
('Skylar', 'Olenchenko'),
('Ashley', 'Ruckledge'),
('Jock', 'Diack'),
('Audrye', 'Bucky'),
('Tabina', 'Colenutt'),
('Arleta', 'Chisholm'),
('Benedict', 'Bindley'),
('Brody', 'Lambrick'),
('Paxton', 'Pierse'),
('Linette', 'Cowdroy'),
('Mikey', 'Geharke'),
('Stacy', 'Insley'),
('Alecia', 'Windrus'),
('Deidre', 'Sudron'),
('Mano', 'Bunce'),
('Nissy', 'Nunnery'),
('Uriah', 'Byrcher'),
('Fergus', 'Glyssanne'),
('Mufi', 'Inchley'),
('Rica', 'Yushin'),
('Gib', 'Poveleye'),
('Oneida', 'Northey'),
('Gian', 'Laugherane'),
('Hinze', 'Johanning'),
('Veronike', 'Snar'),
('Kim', 'Hastin'),
('Keriann', 'Gideon'),
('Cortney', 'Colliver'),
('Rubin', 'Hancorn'),
('Marline', 'Stracey'),
('Muire', 'Sparling'),
('Karin', 'Brack'),
('Edyth', 'Kalisz'),
('Reina', 'Duffree'),
('Inez', 'Quenell'),
('Marvin', 'Iskow'),
('Jana', 'Harkus'),
('Bill', 'Danaher'),
('Lyndsie', 'Shellcross'),
('Camellia', 'Gresswood'),
('Deane', 'Worg'),
('Niles', 'Feitosa'),
('Mart', 'Heinz'),
('Riannon', 'Oxnam'),
('Marsha', 'Darley'),
('Artemus', 'Bauer'),
('Renado', 'Poundsford'),
('Brian', 'Seer'),
('Loretta', 'Stearn'),
('Ned', 'Gosdin'),
('Karylin', 'Crasford'),
('Eleen', 'Swanwick'),
('Cookie', 'Joseland'),
('Enrika', 'Collins'),
('Trix', 'De''Ath'),
('Oran', 'Shury'),
('Delainey', 'Teasell'),
('Gaylord', 'Quigley'),
('Holly-anne', 'Van Arsdall'),
('Rutledge', 'Mougeot'),
('Colin', 'Filgate'),
('Tannie', 'MacMaykin'),
('Trude', 'Joseland'),
('Morissa', 'Coulter'),
('Russell', 'Boullin'),
('Haleigh', 'Bane'),
('Vaughan', 'O''Gavin');

-- create books table
CREATE TABLE IF NOT EXISTS books (
	id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  quantity INT NOT NULL
);

-- insert 100 classic and bestselling books with random quantities (data from claude)
INSERT INTO books (title, quantity) VALUES
    ('To Kill a Mockingbird', floor(random() * 100 + 1)),
    ('1984', floor(random() * 100 + 1)),
    ('Pride and Prejudice', floor(random() * 100 + 1)),
    ('The Great Gatsby', floor(random() * 100 + 1)),
    ('One Hundred Years of Solitude', floor(random() * 100 + 1)),
    ('The Catcher in the Rye', floor(random() * 100 + 1)),
    ('Lord of the Flies', floor(random() * 100 + 1)),
    ('The Hobbit', floor(random() * 100 + 1)),
    ('The Lord of the Rings', floor(random() * 100 + 1)),
    ('Brave New World', floor(random() * 100 + 1)),
    ('The Chronicles of Narnia', floor(random() * 100 + 1)),
    ('The Grapes of Wrath', floor(random() * 100 + 1)),
    ('Jane Eyre', floor(random() * 100 + 1)),
    ('Wuthering Heights', floor(random() * 100 + 1)),
    ('Don Quixote', floor(random() * 100 + 1)),
    ('Moby Dick', floor(random() * 100 + 1)),
    ('The Odyssey', floor(random() * 100 + 1)),
    ('The Divine Comedy', floor(random() * 100 + 1)),
    ('Crime and Punishment', floor(random() * 100 + 1)),
    ('The Brothers Karamazov', floor(random() * 100 + 1)),
    ('War and Peace', floor(random() * 100 + 1)),
    ('Anna Karenina', floor(random() * 100 + 1)),
    ('Les Misérables', floor(random() * 100 + 1)),
    ('The Count of Monte Cristo', floor(random() * 100 + 1)),
    ('The Picture of Dorian Gray', floor(random() * 100 + 1)),
    ('Frankenstein', floor(random() * 100 + 1)),
    ('Dracula', floor(random() * 100 + 1)),
    ('The Adventures of Huckleberry Finn', floor(random() * 100 + 1)),
    ('The Adventures of Tom Sawyer', floor(random() * 100 + 1)),
    ('Little Women', floor(random() * 100 + 1)),
    ('Alice''s Adventures in Wonderland', floor(random() * 100 + 1)),
    ('The Wind in the Willows', floor(random() * 100 + 1)),
    ('Peter Pan', floor(random() * 100 + 1)),
    ('The Secret Garden', floor(random() * 100 + 1)),
    ('Anne of Green Gables', floor(random() * 100 + 1)),
    ('The Call of the Wild', floor(random() * 100 + 1)),
    ('White Fang', floor(random() * 100 + 1)),
    ('The Old Man and the Sea', floor(random() * 100 + 1)),
    ('For Whom the Bell Tolls', floor(random() * 100 + 1)),
    ('The Sun Also Rises', floor(random() * 100 + 1)),
    ('Of Mice and Men', floor(random() * 100 + 1)),
    ('East of Eden', floor(random() * 100 + 1)),
    ('The Pearl', floor(random() * 100 + 1)),
    ('The Sound and the Fury', floor(random() * 100 + 1)),
    ('As I Lay Dying', floor(random() * 100 + 1)),
    ('Light in August', floor(random() * 100 + 1)),
    ('Heart of Darkness', floor(random() * 100 + 1)),
    ('Lord Jim', floor(random() * 100 + 1)),
    ('A Tale of Two Cities', floor(random() * 100 + 1)),
    ('Great Expectations', floor(random() * 100 + 1)),
    ('David Copperfield', floor(random() * 100 + 1)),
    ('Oliver Twist', floor(random() * 100 + 1)),
    ('The Scarlet Letter', floor(random() * 100 + 1)),
    ('The House of the Seven Gables', floor(random() * 100 + 1)),
    ('Uncle Tom''s Cabin', floor(random() * 100 + 1)),
    ('The Red Badge of Courage', floor(random() * 100 + 1)),
    ('The Age of Innocence', floor(random() * 100 + 1)),
    ('The House of Mirth', floor(random() * 100 + 1)),
    ('Ethan Frome', floor(random() * 100 + 1)),
    ('Sister Carrie', floor(random() * 100 + 1)),
    ('The Jungle', floor(random() * 100 + 1)),
    ('Main Street', floor(random() * 100 + 1)),
    ('Babbitt', floor(random() * 100 + 1)),
    ('The Good Earth', floor(random() * 100 + 1)),
    ('Native Son', floor(random() * 100 + 1)),
    ('The Maltese Falcon', floor(random() * 100 + 1)),
    ('The Big Sleep', floor(random() * 100 + 1)),
    ('The Long Goodbye', floor(random() * 100 + 1)),
    ('In Cold Blood', floor(random() * 100 + 1)),
    ('Breakfast at Tiffany''s', floor(random() * 100 + 1)),
    ('On the Road', floor(random() * 100 + 1)),
    ('The Naked and the Dead', floor(random() * 100 + 1)),
    ('Catch-22', floor(random() * 100 + 1)),
    ('Slaughterhouse-Five', floor(random() * 100 + 1)),
    ('The Bell Jar', floor(random() * 100 + 1)),
    ('One Flew Over the Cuckoo''s Nest', floor(random() * 100 + 1)),
    ('Sophie''s Choice', floor(random() * 100 + 1)),
    ('The Color Purple', floor(random() * 100 + 1)),
    ('Beloved', floor(random() * 100 + 1)),
    ('The Handmaid''s Tale', floor(random() * 100 + 1)),
    ('The Road', floor(random() * 100 + 1)),
    ('No Country for Old Men', floor(random() * 100 + 1)),
    ('All the Pretty Horses', floor(random() * 100 + 1)),
    ('Blood Meridian', floor(random() * 100 + 1)),
    ('The Remains of the Day', floor(random() * 100 + 1)),
    ('Never Let Me Go', floor(random() * 100 + 1)),
    ('Atonement', floor(random() * 100 + 1)),
    ('Cloud Atlas', floor(random() * 100 + 1)),
    ('Life of Pi', floor(random() * 100 + 1)),
    ('The Kite Runner', floor(random() * 100 + 1)),
    ('The Road Less Traveled', floor(random() * 100 + 1)),
    ('The Alchemist', floor(random() * 100 + 1)),
    ('The Giver', floor(random() * 100 + 1)),
    ('Bridge to Terabithia', floor(random() * 100 + 1)),
    ('Number the Stars', floor(random() * 100 + 1)),
    ('Roll of Thunder, Hear My Cry', floor(random() * 100 + 1)),
    ('Where the Red Fern Grows', floor(random() * 100 + 1)),
    ('Island of the Blue Dolphins', floor(random() * 100 + 1)),
    ('The Yearling', floor(random() * 100 + 1)),
    ('Johnny Tremain', floor(random() * 100 + 1));

-- create book_transactions table
CREATE TABLE IF NOT EXISTS book_transactions (
	id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users NOT NULL,
  book_id INT REFERENCES books NOT NULL,
  due_date DATE NOT NULL,
  borrowed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  returned_at TIMESTAMP
);


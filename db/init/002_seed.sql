INSERT INTO users (id, name, age, profession, interest_areas)
VALUES
  ('user-ana', 'Ana Lima', 28, 'Software Engineer', ARRAY['technology', 'fiction', 'career']),
  ('user-bruno', 'Bruno Souza', 34, 'Teacher', ARRAY['history', 'education', 'biography']),
  ('user-camila', 'Camila Rocha', 25, 'Designer', ARRAY['art', 'design', 'creativity'])
ON CONFLICT (id) DO NOTHING;

INSERT INTO books (id, title, author, category, subject, area, description)
VALUES
  ('book-clean-arch', 'Clean Architecture', 'Robert C. Martin', 'technology', 'software architecture', 'career', 'Guide to software architecture and code organization.'),
  ('book-practical-ai', 'Practical AI for Developers', 'Sample Author', 'technology', 'artificial intelligence', 'technology', 'Hands-on AI introduction for software developers.'),
  ('book-brazil-history', 'History of Brazil', 'Sample Author', 'history', 'brazilian history', 'history', 'Historical overview of Brazil.'),
  ('book-everyday-design', 'The Design of Everyday Things', 'Don Norman', 'design', 'ux', 'art', 'Core concepts of user-centered design.'),
  ('book-bio-jobs', 'Steve Jobs', 'Walter Isaacson', 'biography', 'leadership', 'biography', 'Biography of Steve Jobs.')
ON CONFLICT (id) DO NOTHING;

INSERT INTO purchases (id, user_id, book_id, quantity)
VALUES
  ('purchase-ana-1', 'user-ana', 'book-clean-arch', 1),
  ('purchase-ana-2', 'user-ana', 'book-practical-ai', 1),
  ('purchase-bruno-1', 'user-bruno', 'book-brazil-history', 1),
  ('purchase-camila-1', 'user-camila', 'book-everyday-design', 1)
ON CONFLICT (id) DO NOTHING;

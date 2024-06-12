<?php
$articles = json_decode(file_get_contents('./articles.json'), true);

$dns = getenv('DB_URL');
$user = 'root';
$pwd = getenv('MYSQL_ROOT_PASSWORD');

$pdo = new PDO($dns, $user, $pwd);

$statement = $pdo->prepare('
  INSERT INTO article (
    title,
    category,
    content,
    image
  ) VALUES (
    :title,
    :category,
    :content,
    :image
)');

foreach ($articles as $article) {
  $statement->bindValue(':title', $article['title']);
  $statement->bindValue(':category', $article['category']);
  $statement->bindValue(':content', $article['content']);
  $statement->bindValue(':image', $article['image']);
  $statement->execute();
}


// declare(strict_types=1);
// error_reporting(E_ALL);
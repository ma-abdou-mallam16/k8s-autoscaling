<?php

$dns = getenv('DB_URL');
$user = 'root';
$pwd = getenv('MYSQL_ROOT_PASSWORD');

try {
    $pdo = new PDO($dns, $user, $pwd, [
        PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
        PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC
    ]);
} catch (PDOException $e) {
    throw new Exception($e->getMessage());
}

return $pdo;

<?php 
// cara membuat object
require_once "data/Person.php";

use Data\Human\Person;

// ini contoh object
$person = new Person("Alber", null);

var_dump($person);

?>
<?php 

namespace Data\Helper;

class Calc{
    var int $angka;
    var int $op;

    function hitung($angka1, $angka2, $op){
        switch ($op){
            case 1:
                echo $angka1 + $angka2;
            break;
            case 2:
                echo $angka1 - $angka2;
            break;
            case 3:
                echo $angka1 * $angka2;
            break;
            case 4:
                echo $angka1 / $angka2;
            break;
        }
    }
}

class Problem{
    
}

?>
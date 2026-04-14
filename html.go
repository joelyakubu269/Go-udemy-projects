<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Unit Converter</title>
</head>
<body>
    <form  method="POST">
        <label for="number">Enter number to be converted</label>
        <input type="number" id="number" name="value" placeholder="100">
        <br>
        <label for="from">From</label>
        
        <select name="from" id="from">
            <option disabled selected>select an option</option>
            <option value="Fahrenheit">Fahrenheit</option>
            <option value="Celsius">Celsius</option>
            <option value="Kelvin">Kelvin</option>
        </select>
        <br>
        <label for="To">To</label>
        <select name="to" id="to">
            <option value=""></option>
             <option value="Fahrenheit">Fahrenheit</option>
            <option value="Celsius">Celsius</option>
            <option value="Kelvin">Kelvin</option>
        </select>
        <input type="submit" value="convert">

    </form>
</body>
</html>


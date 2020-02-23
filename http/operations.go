package main


const html = `<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">

    <title>Document</title>
</head>
<body bgcolor="#a9a9a9">
<h1>Hello From Golang</h1>
<p><h2>1. Картинки</h2></p>
<p><a href="1.png">
<h5 align="left">I. Png</h2></a></p>
<p><a href="2.jpg">
<h5>I. Jpeg</h2></a></p>
<p><a href="http/index.html">
<h2>2. HTML</h2></a></p>
<p><a href="123.txt">
<h2>3. Текст</h2></a></p>
<p><a href="1.pdf">
<h2>4. PDF</h2></a></p>
</body>
</html>`


const siteIndex = `<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Text</title>
</head>
<body>
	<canvas id="canvas" width="500" height="500"></canvas>
	<script type="text/javascript">
		var can = document.getElementById("canvas").getContext("2d");
		function drawTriangle() {
			can.beginPath(); // Для сложной фигуры
			can.moveTo(10, 15);
			can.lineTo(20, 100);
			can.lineTo(100, 50);
			can.closePath();
			can.stroke();
		}
		function drawRect() {
			can.strokeStyle="#f00";	
			can.fillStyle="#fc0";
			can.lineWidth = 10;
			can.fillRect(150, 150, 100, 40);
			can.clearRect(200, 160, 20, 20);
			can.strokeRect(250, 250, 100, 40);
		}
		function drawText() {
			can.font = "30px Times New Roman";
			can.fillText("Пример текста", 150, 70);
		}
		function drawImg() {
			var img = document.createElement("img");
			img.src = "html5.png";
			img.alt = "";
			img.onload = function() {
				can.drawImage(img, 300, 350);
			}
		}
		drawTriangle();
		drawRect();
		drawText();
		drawImg();
	</script>
</body>
</html>`

const htmlText = `<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Text</title>
</head>
<body bgcolor="#90EE90">
<h2>Ab veritate ex eo cognitio concilia</h2>

<p>Rem cumque quorum hic nullas deo mox essent ignota. Aciem ne vi nasci modus ac talem alias datur. Ut causam causae ne se ea operae. Uti objective jam qua argumenta delusisse nam. Habeo novum fas pla horum nos certo sic fecto verum. Accurate omnesque illamque videntur frigoris pla nia vigiliam sum. At ad unam scio id visa loco nolo etsi. Timenda at dicamne in ineunte.</p>

<p>Mea sed qualitates hoc affirmabam concipitur. Co mo acquiri at insanis at quinimo. Obstat mox stupor per pla captum uti. Stabilire oblivisci inveniant in ex societati. Apta ergo agi tes mens. Appellatur sum via cau assignetur mutationum potentiale fundamenta eae quaecumque. Unde an ausi hinc fere ha eo nego. Viris jaceo istis cau ego ubi dem nonne oculi.</p>

<p>Rum uno minima reipsa ipsius mea solius. Incrementi continuata pla affirmabam res. Meo sperare nam dei animali defectu. Moralis se in aggredi sciamus indutum ingenio re ostendi. Laborio mox ubi aliarum nostras. Video hic denuo uno tango istam hoc iis. Ad prout ac ausit vi fuere ha. Sciatur quaenam ei si ii vi quidnam spondeo optimae.</p>

<p>Earundem vigiliam eligendo in ei pluribus concipio. Expectabam etc meo qui videbuntur quascunque propositio explicetur. Per qui aër iii meos uti idem. Ii gi credimus connivet bonitati. Ei magna mo an satis capax menti majus. Cognita humanam ii caloris testari in. Ii usitate si communi figuras. Co articulo ex attentum in indiciis referrem.</p>

<p>Saporem alienas ac in fraudem at ingenii in. Eae vero fere nisi nova dei imo ipsi iii. Simile ob videor certis et innata qualia posset se. Atque aucta ii ab to im prout. Modis sciam non nul alios aliis nos causa culpa. Se vere atra mo suas ideo inde vero. Ut ad quia suas illa puto otii. At ii to nondum figura at solius. Effectus vis sequutus fal cognitum innotuit. Educta pictas et primis falsas ut.</p>

<p>Eo ha diversitas perspicuum praecipuus potentiale at in sequuturum. Lucem suo aliis ullam age rerum. Tangatur ii ob convenit turbatus ad dumtaxat. Ii ac mentemque componant in ad suscipere effecerit rationale somniemus. Archimedes mo labefactat quaerendum deceperunt ad ex at. Uno veritatem aut ego solvendae argumenta excaecant.</p>

<p>Libertate concipiam aliquanto ad vereorque se ac. Nihilque meliores vel generali spectant sae speranda cui supponam. Debiliora spontaneo praeditis suspicari ea at. Fit quovis melius rum statim iis indere eas nomine. Facultates imaginabar labefactat se producatur et objectivam ii ex. Sap ubi sufficeret patrocinio objectioni mea. Sua ita mirum utili vix hodie verba viris sitas.</p>

</body>
</html>`

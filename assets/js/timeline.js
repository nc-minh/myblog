jQuery(document).ready(function ($) {

    const bArray = [];
    const sArray = [4, 6, 8, 10, 12, 14, 16];

    for (let i = 0; i < $('#bubble').width(); i++) {
        bArray.push(i);
    }

    function randomValue(arr) {
        return arr[Math.floor(Math.random() * arr.length)];
    }

    setInterval(function () {
        const size = randomValue(sArray);
        $('#bubble').append('<div class="individual-bubble" style="left: ' + randomValue(bArray) + 'px; width: ' + size + 'px; height:' + size + 'px;"></div>');

        $('.individual-bubble').animate({
            'bottom': '100%',
            'opacity': '-=0.7'
        }, 3000, function () {
            $(this).remove()
        }
        );


    }, 350);

});
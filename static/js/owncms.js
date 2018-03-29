$(function () {

    $('.post_list .post_content img').click(function () {
        $('.preview_img').remove();
        var tmpimg = new Image();
        tmpimg.src = $(this).attr('src');
        ow = tmpimg.width;
        oh = tmpimg.height;
        $('body').append('<img src = "' + $(this).attr('src') + '" class= "preview_img" title="点击关闭" />');
        $('.preview_img').height(oh).width(ow).css({ marginLeft: '-' + ow / 2 + 'px', marginTop: parseInt('-' + oh / 2) + $(document).scrollTop() + 'px' }).show('normal');
    })

    $('.preview_img').live('click', function () {
        $(this).hide('normal');
    })
})
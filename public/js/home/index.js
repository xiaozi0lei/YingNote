/* home/index 首页 */

$(function() {
    // 平滑滚动
    $(".smooth-scroll").click(function(e) {
        e.preventDefault();
        var t = $(this).data("target");
        var targetOffset = $(t).offset().top - 80;
        $('html,body').animate({scrollTop: targetOffset}, 300);
    });

    // 主页滚动图片
    function slider(webImgs, descT) {
        var webImgsLen = webImgs.length;
        var curIndex = 0;
        setInterval(function() {
            webImgs.eq(curIndex).stop().animate({opacity: '0'}, 1000);
            curIndex = (curIndex+1)%webImgsLen;
            var curImg = webImgs.eq(curIndex);
            curImg.stop().animate({opacity: '1'}, 1000);
            descT.text(curImg.data("text"));
        }, 5000);
    }

    slider($(".web-slider"), $("#webText"));
    slider($(".mobile-slider"), $("#mobileText"));

    // 设置 cookie 失效时间
    function setCookie(name, value) {
        var Days = 1;
        var exp  = new Date();
        exp.setTime(exp.getTime() + Days*24*60*60*1000);
        document.cookie = name + "="+ escape (value) + ";expires=" + exp.toGMTString();
    }

    // home 主页右上角的语言切换链接
    $('#lang').find('a').click(function() {
        // 获取所选语言的值
        var lang = $(this).data('lang');
        // 调用上方的函数，设置 cookie
        setCookie('YINGNOTE_LANG', lang);
        // 调用 window 的 location.reload 方法重新加载页面，使用新的语言
        window.location.reload();
    });
});

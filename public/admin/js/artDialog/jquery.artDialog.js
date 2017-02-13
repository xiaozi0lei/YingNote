/*!
 * artDialog 4.1.7
 * Date: 2013-03-03 08:04
 * http://code.google.com/p/artdialog/
 * (c) 2009-2012 TangBin, http://www.planeArt.cn
 *
 * This is licensed under the GNU LGPL, version 2.1 or later.
 * For details, see: http://creativecommons.org/licenses/LGPL/2.1/
 */
(function(e,t,n){e.noop=e.noop||function(){};var r,i,s,o,u=0,a=e(t),f=e(document),l=e("html"),c=document.documentElement,h=t.VBArray&&!t.XMLHttpRequest,p="createTouch"in document&&!("onmousemove"in c)||/(iPhone|iPad|iPod)/i.test(navigator.userAgent),d="artDialog"+ +(new Date),v=function(t,i,s){t=t||{};if(typeof t=="string"||t.nodeType===1)t={content:t,fixed:!p};var o,a=v.defaults,f=t.follow=this.nodeType===1&&this||t.follow;for(var l in a)t[l]===n&&(t[l]=a[l]);return e.each({ok:"yesFn",cancel:"noFn",close:"closeFn",init:"initFn",okVal:"yesText",cancelVal:"noText"},function(e,r){t[e]=t[e]!==n?t[e]:t[r]}),typeof f=="string"&&(f=e(f)[0]),t.id=f&&f[d+"follow"]||t.id||d+u,o=v.list[t.id],f&&o?o.follow(f).zIndex().focus():o?o.zIndex().focus():(p&&(t.fixed=!1),e.isArray(t.button)||(t.button=t.button?[t.button]:[]),i!==n&&(t.ok=i),s!==n&&(t.cancel=s),t.ok&&t.button.push({name:t.okVal,callback:t.ok,focus:!0}),t.cancel&&t.button.push({name:t.cancelVal,callback:t.cancel}),v.defaults.zIndex=t.zIndex,u++,v.list[t.id]=r?r._init(t):new v.fn._init(t))};v.fn=v.prototype={version:"4.1.7",closed:!0,_init:function(e){var n=this,i,s=e.icon,o=s&&(h?{png:"icons/"+s+".png"}:{backgroundImage:"url('"+e.path+"/skins/icons/"+s+".png')"});return n.closed=!1,n.config=e,n.DOM=i=n.DOM||n._getDOM(),i.wrap.addClass(e.skin),i.close[e.cancel===!1?"hide":"show"](),i.icon[0].style.display=s?"":"none",i.iconBg.css(o||{background:"none"}),i.se.css("cursor",e.resize?"se-resize":"auto"),i.title.css("cursor",e.drag?"move":"auto"),i.content.css("padding",e.padding),n[e.show?"show":"hide"](!0),n.button(e.button).title(e.title).content(e.content,!0).size(e.width,e.height).time(e.time),e.follow?n.follow(e.follow):n.position(e.left,e.top),n.zIndex().focus(),e.lock&&n.lock(),n._addEvent(),n._ie6PngFix(),r=null,e.init&&e.init.call(n,t),n},content:function(e){var t,r,i,s,o=this,u=o.DOM,a=u.wrap[0],f=a.offsetWidth,l=a.offsetHeight,c=parseInt(a.style.left),h=parseInt(a.style.top),p=a.style.width,d=u.content,v=d[0];return o._elemBack&&o._elemBack(),a.style.width="auto",e===n?v:(typeof e=="string"?d.html(e):e&&e.nodeType===1&&(s=e.style.display,t=e.previousSibling,r=e.nextSibling,i=e.parentNode,o._elemBack=function(){t&&t.parentNode?t.parentNode.insertBefore(e,t.nextSibling):r&&r.parentNode?r.parentNode.insertBefore(e,r):i&&i.appendChild(e),e.style.display=s,o._elemBack=null},d.html(""),v.appendChild(e),e.style.display="block"),arguments[1]||(o.config.follow?o.follow(o.config.follow):(f=a.offsetWidth-f,l=a.offsetHeight-l,c-=f/2,h-=l/2,a.style.left=Math.max(c,0)+"px",a.style.top=Math.max(h,0)+"px"),p&&p!=="auto"&&(a.style.width=a.offsetWidth+"px"),o._autoPositionType()),o._ie6SelectFix(),o._runScript(v),o)},title:function(e){var t=this.DOM,r=t.wrap,i=t.title,s="aui_state_noTitle";return e===n?i[0]:(e===!1?(i.hide().html(""),r.addClass(s)):(i.show().html(e||""),r.removeClass(s)),this)},position:function(e,t){var r=this,i=r.config,s=r.DOM.wrap[0],o=h?!1:i.fixed,u=h&&r.config.fixed,l=f.scrollLeft(),c=f.scrollTop(),p=o?0:l,d=o?0:c,v=a.width(),m=a.height(),g=s.offsetWidth,y=s.offsetHeight,b=s.style;if(e||e===0)r._left=e.toString().indexOf("%")!==-1?e:null,e=r._toNumber(e,v-g),typeof e=="number"?(e=u?e+=l:e+p,b.left=Math.max(e,p)+"px"):typeof e=="string"&&(b.left=e);if(t||t===0)r._top=t.toString().indexOf("%")!==-1?t:null,t=r._toNumber(t,m-y),typeof t=="number"?(t=u?t+=c:t+d,b.top=Math.max(t,d)+"px"):typeof t=="string"&&(b.top=t);return e!==n&&t!==n&&(r._follow=null,r._autoPositionType()),r},size:function(e,t){var n,r,i,s,o=this,u=o.config,f=o.DOM,l=f.wrap,c=f.main,h=l[0].style,p=c[0].style;return e&&(o._width=e.toString().indexOf("%")!==-1?e:null,n=a.width()-l[0].offsetWidth+c[0].offsetWidth,i=o._toNumber(e,n),e=i,typeof e=="number"?(h.width="auto",p.width=Math.max(o.config.minWidth,e)+"px",h.width=l[0].offsetWidth+"px"):typeof e=="string"&&(p.width=e,e==="auto"&&l.css("width","auto"))),t&&(o._height=t.toString().indexOf("%")!==-1?t:null,r=a.height()-l[0].offsetHeight+c[0].offsetHeight,s=o._toNumber(t,r),t=s,typeof t=="number"?p.height=Math.max(o.config.minHeight,t)+"px":typeof t=="string"&&(p.height=t)),o._ie6SelectFix(),o},follow:function(t){var n,r=this,i=r.config;if(typeof t=="string"||t&&t.nodeType===1)n=e(t),t=n[0];if(!t||!t.offsetWidth&&!t.offsetHeight)return r.position(r._left,r._top);var s=d+"follow",o=a.width(),u=a.height(),l=f.scrollLeft(),c=f.scrollTop(),p=n.offset(),v=t.offsetWidth,m=t.offsetHeight,g=h?!1:i.fixed,y=g?p.left-l:p.left,b=g?p.top-c:p.top,w=r.DOM.wrap[0],E=w.style,S=w.offsetWidth,x=w.offsetHeight,T=y-(S-v)/2,N=b+m,C=g?0:l,k=g?0:c;return T=T<C?y:T+S>o&&y-S>C?y-S+v:T,N=N+x>u+k&&b-x>k?b-x:N,E.left=T+"px",E.top=N+"px",r._follow&&r._follow.removeAttribute(s),r._follow=t,t[s]=i.id,r._autoPositionType(),r},button:function(){var t=this,r=arguments,i=t.DOM,s=i.buttons,o=s[0],u="aui_state_highlight",a=t._listeners=t._listeners||{},f=e.isArray(r[0])?r[0]:[].slice.call(r);return r[0]===n?o:(e.each(f,function(n,r){var i=r.name,s=!a[i],f=s?document.createElement("button"):a[i].elem;a[i]||(a[i]={}),r.callback&&(a[i].callback=r.callback),r.className&&(f.className=r.className),r.focus&&(t._focus&&t._focus.removeClass(u),t._focus=e(f).addClass(u),t.focus()),f.setAttribute("type","button"),f[d+"callback"]=i,f.disabled=!!r.disabled,s&&(f.innerHTML=i,a[i].elem=f,o.appendChild(f))}),s[0].style.display=f.length?"":"none",t._ie6SelectFix(),t)},show:function(){return this.DOM.wrap.show(),!arguments[0]&&this._lockMaskWrap&&this._lockMaskWrap.show(),this},hide:function(){return this.DOM.wrap.hide(),!arguments[0]&&this._lockMaskWrap&&this._lockMaskWrap.hide(),this},close:function(){if(this.closed)return this;var e=this,n=e.DOM,i=n.wrap,s=v.list,o=e.config.close,u=e.config.follow;e.time();if(typeof o=="function"&&o.call(e,t)===!1)return e;e.unlock(),e._elemBack&&e._elemBack(),i[0].className=i[0].style.cssText="",n.title.html(""),n.content.html(""),n.buttons.html(""),v.focus===e&&(v.focus=null),u&&u.removeAttribute(d+"follow"),delete s[e.config.id],e._removeEvent(),e.hide(!0)._setAbsolute();for(var a in e)e.hasOwnProperty(a)&&a!=="DOM"&&delete e[a];return r?i.remove():r=e,e},time:function(e){var t=this,n=t.config.cancelVal,r=t._timer;return r&&clearTimeout(r),e&&(t._timer=setTimeout(function(){t._click(n)},1e3*e)),t},focus:function(){try{if(this.config.focus){var e=this._focus&&this._focus[0]||this.DOM.close[0];e&&e.focus()}}catch(t){}return this},zIndex:function(){var e=this,t=e.DOM,n=t.wrap,r=v.focus,i=v.defaults.zIndex++;return n.css("zIndex",i),e._lockMask&&e._lockMask.css("zIndex",i-1),r&&r.DOM.wrap.removeClass("aui_state_focus"),v.focus=e,n.addClass("aui_state_focus"),e},lock:function(){if(this._lock)return this;var t=this,n=v.defaults.zIndex-1,r=t.DOM.wrap,i=t.config,s=f.width(),o=f.height(),u=t._lockMaskWrap||e(document.body.appendChild(document.createElement("div"))),a=t._lockMask||e(u[0].appendChild(document.createElement("div"))),l="(document).documentElement",c=p?"width:"+s+"px;height:"+o+"px":"width:100%;height:100%",d=h?"position:absolute;left:expression("+l+".scrollLeft);top:expression("+l+".scrollTop);width:expression("+l+".clientWidth);height:expression("+l+".clientHeight)":"";return t.zIndex(),r.addClass("aui_state_lock"),u[0].style.cssText=c+";position:fixed;z-index:"+n+";top:0;left:0;overflow:hidden;"+d,a[0].style.cssText="height:100%;background:"+i.background+";filter:alpha(opacity=0);opacity:0",h&&a.html('<iframe src="about:blank" style="width:100%;height:100%;position:absolute;top:0;left:0;z-index:-1;filter:alpha(opacity=0)"></iframe>'),a.stop(),a.bind("click",function(){t._reset()}).bind("dblclick",function(){t._click(t.config.cancelVal)}),i.duration===0?a.css({opacity:i.opacity}):a.animate({opacity:i.opacity},i.duration),t._lockMaskWrap=u,t._lockMask=a,t._lock=!0,t},unlock:function(){var e=this,t=e._lockMaskWrap,n=e._lockMask;if(!e._lock)return e;var i=t[0].style,s=function(){h&&(i.removeExpression("width"),i.removeExpression("height"),i.removeExpression("left"),i.removeExpression("top")),i.cssText="display:none",r&&t.remove()};return n.stop().unbind(),e.DOM.wrap.removeClass("aui_state_lock"),e.config.duration?n.animate({opacity:0},e.config.duration,s):s(),e._lock=!1,e},_getDOM:function(){var t=document.createElement("div"),n=document.body;t.style.cssText="position:absolute;left:0;top:0",t.innerHTML=v._templates,n.insertBefore(t,n.firstChild);var r,i=0,s={wrap:e(t)},o=t.getElementsByTagName("*"),u=o.length;for(;i<u;i++)r=o[i].className.split("aui_")[1],r&&(s[r]=e(o[i]));return s},_toNumber:function(e,t){if(!e&&e!==0||typeof e=="number")return e;var n=e.length-1;return e.lastIndexOf("px")===n?e=parseInt(e):e.lastIndexOf("%")===n&&(e=parseInt(t*e.split("%")[0]/100)),e},_ie6PngFix:h?function(){var e=0,t,n,r,i,s=v.defaults.path+"/skins/",o=this.DOM.wrap[0].getElementsByTagName("*");for(;e<o.length;e++)t=o[e],n=t.currentStyle.png,n&&(r=s+n,i=t.runtimeStyle,i.backgroundImage="none",i.filter="progid:DXImageTransform.Microsoft.AlphaImageLoader(src='"+r+"',sizingMethod='crop')")}:e.noop,_ie6SelectFix:h?function(){var e=this.DOM.wrap,t=e[0],n=d+"iframeMask",r=e[n],i=t.offsetWidth,s=t.offsetHeight;i+="px",s+="px",r?(r.style.width=i,r.style.height=s):(r=t.appendChild(document.createElement("iframe")),e[n]=r,r.src="about:blank",r.style.cssText="position:absolute;z-index:-1;left:0;top:0;filter:alpha(opacity=0);width:"+i+";height:"+s)}:e.noop,_runScript:function(e){var t,n=0,r=0,i=e.getElementsByTagName("script"),s=i.length,o=[];for(;n<s;n++)i[n].type==="text/dialog"&&(o[r]=i[n].innerHTML,r++);o.length&&(o=o.join(""),t=new Function(o),t.call(this))},_autoPositionType:function(){this[this.config.fixed?"_setFixed":"_setAbsolute"]()},_setFixed:function(){return h&&e(function(){var t="backgroundAttachment";l.css(t)!=="fixed"&&e("body").css(t)!=="fixed"&&l.css({zoom:1,backgroundImage:"url(about:blank)",backgroundAttachment:"fixed"})}),function(){var e=this.DOM.wrap,t=e[0].style;if(h){var n=parseInt(e.css("left")),r=parseInt(e.css("top")),i=f.scrollLeft(),s=f.scrollTop(),o="(document.documentElement)";this._setAbsolute(),t.setExpression("left","eval("+o+".scrollLeft + "+(n-i)+') + "px"'),t.setExpression("top","eval("+o+".scrollTop + "+(r-s)+') + "px"')}else t.position="fixed"}}(),_setAbsolute:function(){var e=this.DOM.wrap[0].style;h&&(e.removeExpression("left"),e.removeExpression("top")),e.position="absolute"},_click:function(e){var n=this,r=n._listeners[e]&&n._listeners[e].callback;return typeof r!="function"||r.call(n,t)!==!1?n.close():n},_reset:function(e){var t,n=this,r=n._winSize||a.width()*a.height(),i=n._follow,s=n._width,o=n._height,u=n._left,f=n._top;if(e){t=n._winSize=a.width()*a.height();if(r===t)return}(s||o)&&n.size(s,o),i?n.follow(i):(u||f)&&n.position(u,f)},_addEvent:function(){var e,n=this,r=n.config,i="CollectGarbage"in t,s=n.DOM;n._winResize=function(){e&&clearTimeout(e),e=setTimeout(function(){n._reset(i)},40)},a.bind("resize",n._winResize),s.wrap.bind("click",function(e){var t=e.target,i;if(t.disabled)return!1;if(t===s.close[0])return n._click(r.cancelVal),!1;i=t[d+"callback"],i&&n._click(i),n._ie6SelectFix()}).bind("mousedown",function(){n.zIndex()})},_removeEvent:function(){var e=this,t=e.DOM;t.wrap.unbind(),a.unbind("resize",e._winResize)}},v.fn._init.prototype=v.fn,e.fn.dialog=e.fn.artDialog=function(){var e=arguments;return this[this.live?"live":"bind"]("click",function(){return v.apply(this,e),!1}),this},v.focus=null,v.get=function(e){return e===n?v.list:v.list[e]},v.list={},f.bind("keydown",function(e){var t=e.target,n=t.nodeName,r=/^INPUT|TEXTAREA$/,i=v.focus,s=e.keyCode;if(!i||!i.config.esc||r.test(n))return;s===27&&i._click(i.config.cancelVal)}),o=t._artDialog_path||function(e,t,n){for(t in e)e[t].src&&e[t].src.indexOf("artDialog")!==-1&&(n=e[t]);return i=n||e[e.length-1],n=i.src.replace(/\\/g,"/"),n.lastIndexOf("/")<0?".":n.substring(0,n.lastIndexOf("/"))}
// skin life
(document.getElementsByTagName("script")),
s="blue"; /*i.src.split("skin=")[1]*/;if(s){var m=document.createElement("link");m.rel="stylesheet",m.href=o+"/skins/"+s+".css?"+v.fn.version,i.parentNode.insertBefore(m,i)}a.bind("load",function(){setTimeout(function(){if(u)return;v({left:"-9999em",time:9,fixed:!1,lock:!1,focus:!1})},150)});try{document.execCommand("BackgroundImageCache",!1,!0)}catch(g){}v._templates='<div class="aui_outer"><table class="aui_border"><tbody><tr><td class="aui_nw"></td><td class="aui_n"></td><td class="aui_ne"></td></tr><tr><td class="aui_w"></td><td class="aui_c"><div class="aui_inner"><table class="aui_dialog"><tbody><tr><td colspan="2" class="aui_header"><div class="aui_titleBar"><div class="aui_title"></div><a class="aui_close" href="javascript:/*artDialog*/;">\u00d7</a></div></td></tr><tr><td class="aui_icon"><div class="aui_iconBg"></div></td><td class="aui_main"><div class="aui_content"></div></td></tr><tr><td colspan="2" class="aui_footer"><div class="aui_buttons"></div></td></tr></tbody></table></div></td><td class="aui_e"></td></tr><tr><td class="aui_sw"></td><td class="aui_s"></td><td class="aui_se"></td></tr></tbody></table></div>',v.defaults={content:'<div class="aui_loading"><span>loading..</span></div>',title:"\u6d88\u606f",button:null,ok:null,cancel:null,init:null,close:null,okVal:"\u786e\u5b9a",cancelVal:"\u53d6\u6d88",width:"auto",height:"auto",minWidth:96,minHeight:32,padding:"20px 25px",skin:"",icon:null,time:null,esc:!0,focus:!0,show:!0,follow:null,path:o,lock:!1,background:"#000",opacity:.7,duration:300,fixed:!1,left:"50%",top:"38.2%",zIndex:1987,resize:!0,drag:!0},t.artDialog=e.dialog=e.artDialog=v})(this.art||this.jQuery&&(this.art=jQuery),this),function(e){var t,n,r=e(window),i=e(document),s=document.documentElement,o=!("minWidth"in s.style),u="onlosecapture"in s,a="setCapture"in s;artDialog.dragEvent=function(){var e=this,t=function(t){var n=e[t];e[t]=function(){return n.apply(e,arguments)}};t("start"),t("move"),t("end")},artDialog.dragEvent.prototype={onstart:e.noop,start:function(e){return i.bind("mousemove",this.move).bind("mouseup",this.end),this._sClientX=e.clientX,this._sClientY=e.clientY,this.onstart(e.clientX,e.clientY),!1},onmove:e.noop,move:function(e){return this._mClientX=e.clientX,this._mClientY=e.clientY,this.onmove(e.clientX-this._sClientX,e.clientY-this._sClientY),!1},onend:e.noop,end:function(e){return i.unbind("mousemove",this.move).unbind("mouseup",this.end),this.onend(e.clientX,e.clientY),!1}},n=function(e){var n,s,f,l,c,h,p=artDialog.focus,d=p.DOM,v=d.wrap,m=d.title,g=d.main,y="getSelection"in window?function(){window.getSelection().removeAllRanges()}:function(){try{document.selection.empty()}catch(e){}};t.onstart=function(e,n){h?(s=g[0].offsetWidth,f=g[0].offsetHeight):(l=v[0].offsetLeft,c=v[0].offsetTop),i.bind("dblclick",t.end),!o&&u?m.bind("losecapture",t.end):r.bind("blur",t.end),a&&m[0].setCapture(),v.addClass("aui_state_drag"),p.focus()},t.onmove=function(e,t){if(h){var r=v[0].style,i=g[0].style,o=e+s,u=t+f;r.width="auto",i.width=Math.max(0,o)+"px",r.width=v[0].offsetWidth+"px",i.height=Math.max(0,u)+"px"}else{var i=v[0].style,a=Math.max(n.minX,Math.min(n.maxX,e+l)),d=Math.max(n.minY,Math.min(n.maxY,t+c));i.left=a+"px",i.top=d+"px"}y(),p._ie6SelectFix()},t.onend=function(e,n){i.unbind("dblclick",t.end),!o&&u?m.unbind("losecapture",t.end):r.unbind("blur",t.end),a&&m[0].releaseCapture(),o&&!p.closed&&p._autoPositionType(),v.removeClass("aui_state_drag")},h=e.target===d.se[0]?!0:!1,n=function(){var e,t,n=p.DOM.wrap[0],s=n.style.position==="fixed",o=n.offsetWidth,u=n.offsetHeight,a=r.width(),f=r.height(),l=s?0:i.scrollLeft(),c=s?0:i.scrollTop(),e=a-o+l;return t=f-u+c,{minX:l,minY:c,maxX:e,maxY:t}}(),t.start(e)},i.bind("mousedown",function(e){var r=artDialog.focus;if(!r)return;var i=e.target,s=r.config,o=r.DOM;if(s.drag!==!1&&i===o.title[0]||s.resize!==!1&&i===o.se[0])return t=t||new artDialog.dragEvent,n(e),!1})}(this.art||this.jQuery&&(this.art=jQuery))
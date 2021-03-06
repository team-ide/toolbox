import server from "@/server/index.js";
import source from "@/source/index.js";
import teamide from 'teamide-ui'
import jQuery from 'jquery'

import md5 from 'js-md5';
let tool = {};
tool.jQuery = jQuery;
tool.md5 = md5;
for (let key in teamide) {
    tool[key] = teamide[key];
}

tool.open = function () {
    server.open().then(res => {
        let value = res.value;
        tool.setSource(value);
        if (value) {
            tool.initSession();
        }
    }).catch(e => {
        tool.setSource(null);
    });
};

tool.initSession = function () {
    server.session().then(res => {
        let value = res.value;
        tool.setSession(value);
    }).catch(e => {
        tool.setSession(null);
    });
};

tool.setSource = function (value) {
    if (value != null) {
        source.TIMESTAMP = value.TIMESTAMP;
        source.enum_map = value.enum_map || {};
        server.initWorkers(value.workers);

        value.workers = value.workers || [];
        value.workers.forEach(worker => {
            source[worker.name] = { configs: worker.configs };
        });
    } else {
        source.served = false;
        source.inited = true;
    }
};

tool.setSession = function (value) {
    value = value || {};
    source.login_user = value.user;
    source.served = true;
    source.inited = true;
};

tool.doLogout = function () {
    server.logout().then(res => {
        tool.initSession();
    }).catch(e => {
    });
};
tool.replaceAll = function (str, s1, s2) {
    if (str == null) {
        return str;
    }
    str = '' + str;
    return str.replace(new RegExp(s1, "gm"), s2);
};

tool.computeFontSize = function (str, size, family) {
    let spanDom = document.createElement("span");
    spanDom.style.fontSize = size;
    spanDom.style.opacity = "0";
    // spanDom.style.fontFamily = family;
    if (tool.isNotEmpty(str)) {
        str = tool.replaceAll(str, " ", "&nbsp;")
        str = tool.replaceAll(str, "<", "&lt;")
        str = tool.replaceAll(str, ">", "&gt;")
    }
    spanDom.innerHTML = str;
    document.body.append(spanDom);
    let sizeD = {};
    sizeD.width = spanDom.offsetWidth;
    sizeD.height = spanDom.offsetHeight;
    spanDom.remove();
    return sizeD;
};

/**
 * ????????????div?????????
 */
tool.getElementTop = (e) => {
    var offset = e.offsetTop
    if (e.offsetParent != null) {
        offset += tool.getElementTop(e.offsetParent)
    }
    return offset
};
/**
 * ????????????div?????????
 */
tool.getElementLeft = (e) => {
    var offset = e.offsetLeft
    if (e.offsetParent != null) {
        offset += tool.getElementLeft(e.offsetParent)
    }
    return offset
};
tool.getElementBottom = (e) => {
    var wh = tool.jQuery(window).height(),//?????????????????????
        ot = tool.jQuery(e).offset().top,//???????????????????????????
        ds = tool.jQuery(document.documentElement).scrollTop(),//??????????????????// ds = window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop; 
        icoimg_h = tool.jQuery(e).height();//???????????????

    // bh+$("div").height()+[$("div").offset().top-$(document).scrollTop()]=$(window).height();

    var bh = wh - icoimg_h - [ot - ds];

    return bh
};

tool.change = (obj, name, value) => {
    obj = obj || {};
    obj[name] = value;
};
tool.changeForMap = (obj, map) => {
    obj = obj || {};
    for (let name in map) {
        obj[name] = map[name];
    }

};
tool.push = (parent, name, value) => {
    parent = parent || {};
    parent[name] = parent[name] || [];
    value = value || {};
    parent[name].push(value);
};
tool.del = (parent, name, value) => {
    parent = parent || {};
    parent[name] = parent[name] || [];
    value = value || {};
    let index = parent[name].indexOf(value);
    if (index >= 0) {
        parent[name].splice(index, 1);
    }
};
tool.up = (parent, name, value) => {
    parent = parent || {};
    parent[name] = parent[name] || [];
    value = value || {};
    let index = parent[name].indexOf(value);
    if (index > 0) {
        parent[name].splice(index, 1);
        parent[name].splice(index - 1, 0, value);
    }
};
tool.down = (parent, name, value) => {
    parent = parent || {};
    parent[name] = parent[name] || [];
    value = value || {};
    let index = parent[name].indexOf(value);
    if (index >= 0 && index < parent[name].length - 1) {
        parent[name].splice(index, 1);
        parent[name].splice(index + 1, 0, value);
    }
};
tool.initInputWidth = (event, options) => {
    let target = tool.jQuery(event.target || event);
    let value = target.val();
    let str = value;
    let w = 10;
    if (target.find('option').length > 0) {
        w = 30;
        target.find('option').each((index, one) => {
            one = tool.jQuery(one);
            if (one.attr('value') == value) {
                str = one.attr('text') || one.text();
            }
        });
    }
    let compute = tool.computeFontSize(str, "12px");
    tool.jQuery(target).css('width', (compute.width + w) + 'px')
};
tool.getEnum = (type, value) => {
    let result = null;
    let options = source.enum_map[type];
    if (options) {
        options.forEach(one => {
            if (one.value == value) {
                result = one;
            }
        });
        if (result == null) {
            options.forEach(one => {
                if (one.text == value) {
                    result = one;
                }
            });
        }
    }
    return result;
};
tool.getEnumName = (type, value) => {
    let one = tool.getEnum(type, value);
    if (one == null) {
        return;
    }
    return one.name;
};
tool.getEnumValue = (type, value) => {
    let one = tool.getEnum(type, value);
    if (one == null) {
        return;
    }
    return one.value;
};
tool.getEnumText = (type, value) => {
    let one = tool.getEnum(type, value);
    if (one == null) {
        return;
    }
    return one.text;
};
tool.getEnumTextSpan = (type, value) => {
    let text = '';
    let color = '';
    let one = tool.getEnum(type, value);
    if (one) {
        text = one.text;
        color = one.color;
    }
    let html = '';
    html += '<span '
    if (tool.isNotEmpty(color)) {
        if (color.startsWith('#')) {
            html += ' style="color:' + color + ';" ';
        } else {
            html += ' class="color-' + color + '" ';
        }
    }
    html += ' >';
    html += text;
    html += '</span>'
    return html;
};

let chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789".split('');
tool.getRandom = function (length) {
    let res = '';
    let max = chars.length;
    let min = 0;
    for (let i = 0; i < length; i++) {
        let index = Math.floor(Math.random() * (max - min) + min);
        res += chars[index];
    }
    return res;
};
tool.formatDateStr = function (datetime) {
    var value = datetime || "";
    if (!teamide.isEmpty(value)) {
        value = '' + value;
        value = value.replace(/-/g, '');
        value = value.replace(/:/g, '');
        value = value.replace(/ /g, '');
        value = value.replace(/\//g, '');
        if (value.length == 4) {
            value = value.substring(0, 2) + ':' + value.substring(2, 4);
        } else if (value.length == 8) {
            value = value.substring(0, 4) + '-' + value.substring(4, 6) + '-' + value.substring(6, 8);
        } else if (value.length >= 12) {
            value = value.substring(0, 4) + '-' + value.substring(4, 6) + '-' + value.substring(6, 8) + ' ' + value.substring(8, 10) + ':' + value.substring(10, 12);
        }
    }

    return value;
};


tool.getCacheKey = function (key) {
    return source.ROOT_URL + '-' + key;
};

tool.getCache = function (key) {
    key = tool.getCacheKey(key);
    return localStorage.getItem(key);
};
tool.setCache = function (key, value) {
    key = tool.getCacheKey(key);
    localStorage.setItem(key, value);
};
tool.removeCache = function (key) {
    key = tool.getCacheKey(key);
    return localStorage.removeItem(key);
};

tool.initClientUUID = function () {
    let key = 'TEAM_IDE_CLIENT_UUID';
    if (tool.isEmpty(source.CLIENT_UUID)) {
        source.CLIENT_UUID = tool.getCache(key);
        if (tool.isEmpty(source.CLIENT_UUID) || source.CLIENT_UUID.length < 32) {
            source.CLIENT_UUID = tool.getRandom(32);
            tool.setCache(key, source.CLIENT_UUID);
        }
    } else {
        tool.setCache(key, source.CLIENT_UUID);
    }
};

tool.isTrimEmpty = function (arg) {
    if (arg == null) {
        return true;
    }
    return tool.isEmpty(('' + arg).trim());
};

tool.isNotTrimEmpty = function (arg) {
    return !tool.isTrimEmpty(arg);
};

tool.initTreeWidth = function (tree, treeBox) {

    if (tree.initWidthIng) {
        return;
    }
    tree.initWidthIng = true;

    tree.$nextTick(() => {
        let hasActive = false;
        if (
            tree.$el.getElementsByClassName("v-enter-active").length > 0 ||
            tree.$el.getElementsByClassName("v-leave-active").length > 0
        ) {
            hasActive = true;
        } else { }
        if (hasActive) {
            window.setTimeout(() => {
                tree.initWidthIng = false;
                tool.initTreeWidth(tree, treeBox);
            }, 100);
        } else {
            tree.$el.style.minWidth = "0px";
            let minWidth = 0;
            if (treeBox) {
                if (minWidth < treeBox.scrollWidth) {
                    minWidth = treeBox.scrollWidth;
                }
            }
            tree.minWidth = minWidth;
            tree.$el.style.minWidth = minWidth + "px";
            tree.initWidthIng = false;
        }
    });
};

source.CLIENT_TAB_KEY = tool.getRandom(32);
export default tool;
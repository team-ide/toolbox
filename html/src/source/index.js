let source = {};

source.login_user = null;
source.served = false;
source.inited = false;
source.enum_map = null;

source.contextmenu = {
    menus: null
};

source.ROOT_URL = location.origin + location.pathname;
if (!source.ROOT_URL.endsWith('/')) {
    source.ROOT_URL = source.ROOT_URL + '/';
}
source.UPLOAD_URL = source.ROOT_URL;
source.FILE_URL = source.ROOT_URL;


export default source;
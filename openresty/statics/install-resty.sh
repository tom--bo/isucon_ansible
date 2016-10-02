#!/bin/sh
cd /tmp/openresty-1.11.2.1
export PATH="/sbin/:$PATH"
./configure --prefix=/opt/openresty --with-http_realip_module --with-http_addition_module --with-http_sub_module --with-http_dav_module --with-http_flv_module --with-http_gzip_static_module --with-http_auth_request_module --with-http_random_index_module --with-http_secure_link_module --with-http_degradation_module --with-http_stub_status_module
sudo make install

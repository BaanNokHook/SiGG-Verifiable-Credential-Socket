/*
* Go-lang-Vcs-NEXTClan 
*/

\! echo "Configuring MySQL users...";

/*
oidc provider (hydra)
*/
CREATE USER 'thirdpartyoidc'@'%' IDENTIFIED BY 'thirdpartyoidc-secret-pw';
CREATE DATABASE thirdpartyoidc;
GRANT ALL PRIVILEGES ON thirdpartyoidc.* TO 'thirdpartyoidc'@'%';

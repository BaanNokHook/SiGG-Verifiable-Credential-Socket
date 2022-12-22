#
#  Go-lang-Vcs-NEXTClan 
#

@all
@vc_dev_api
Feature: Request DID Config
  Scenario Outline: Request DID Config
    When I request did config for "<profileType>" with ID "<profileId>" and type "<credentialType>"
    Then I receive response with status code "200" for didconfig

    Examples:
      | profileType | profileId                 | credentialType |
      | issuer      | i_myprofile_ud_P256k1     | ldp            |
      | issuer      | i_myprofile_ud_es256_jwt  | jwt            |
      | verifier    | v_myprofile_ldp           | ldp            |
      | verifier    | v_myprofile_jwt           | jwt            |
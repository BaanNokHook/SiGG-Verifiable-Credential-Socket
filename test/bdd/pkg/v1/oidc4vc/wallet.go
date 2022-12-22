/*
Go-lang-Vcs-NEXTClan
*/

package oidc4vc

func (s *Steps) saveCredentialsInWallet() error {
	for _, cred := range s.bddContext.CreatedCredentialsSet {
		if err := s.walletRunner.SaveCredentialInWallet(cred); err != nil {
			return err
		}
	}

	return nil
}

func (s *Steps) createWallet() error {
	err := s.walletRunner.CreateWallet()
	if err != nil {
		return err
	}

	s.bddContext.CredentialSubject = s.walletRunner.GetConfig().WalletParams.DidID
	return nil
}

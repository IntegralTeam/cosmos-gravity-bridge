use deep_space::error::PrivateKeyError;
use deep_space::error::{Bip39Error, HdWalletError};
use deep_space::PrivateKey;

pub trait WalletPathSwitcher {
    /// This function is a wrapper of PrivateKey::from_hd_wallet_path.
    /// It takes 3rd additional parameter 'prefix' to determine the path for the HD wallet.
    /// There are default values for 'cosmos' and 'und'
    fn from_phrase_with_prefix(
        phrase: &str,
        passphrase: &str,
        prefix: &str,
    ) -> Result<PrivateKey, PrivateKeyError> {
        let path = match prefix {
            "cosmos" => "m/44'/118'/0'/0/0",
            "und" => "m/44'/5555'/0'/0/0",
            _ => return Err(HdWalletError::InvalidPathSpec("Can not determine path for this prefix".to_string()).into()),
        };

        if phrase.is_empty() {
            return Err(HdWalletError::Bip39Error(Bip39Error::BadWordCount(0)).into());
        }
        PrivateKey::from_hd_wallet_path(path, phrase, passphrase)
    }
}

impl WalletPathSwitcher for PrivateKey {}

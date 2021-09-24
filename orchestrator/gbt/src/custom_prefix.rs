use deep_space::error::PrivateKeyError;
use deep_space::error::{Bip39Error, HdWalletError};
use deep_space::PrivateKey;

static mut CURRENT_PATH: &'static str = "";

pub fn init_cosmos_key_generation(prefix: &str) {
    let path = match prefix {
        "cosmos" => "m/44'/118'/0'/0/0",
        "und" => "m/44'/5555'/0'/0/0",
        _ => {
            unreachable!(format!("No known wallet path for prefix: {}", prefix))
        }
    };
    unsafe {
        CURRENT_PATH = path;
    }
}

pub trait CustomPrefix
where
    Self: Sized,
{
    type Err;

    /// This function is a wrapper of PrivateKey::from_hd_wallet_path.
    /// It uses different HD wallet paths depending on a prefix, provided to
    /// init_cosmos_key_generation function.
    /// There are default values for 'cosmos' and 'und'
    fn from_phrase_with_custom_prefix(
        phrase: &str,
        passphrase: &str,
    ) -> Result<Self, Self::Err>;
}

impl CustomPrefix for PrivateKey {
    type Err = PrivateKeyError;

    fn from_phrase_with_custom_prefix(
        phrase: &str,
        passphrase: &str,
    ) -> Result<Self, Self::Err> {

        let path = unsafe { CURRENT_PATH };
        if path.len() == 0 {
            return Err(HdWalletError::InvalidPathSpec(
                "Cosmos key generation is not initialized".to_string(),
            )
            .into());
        }

        if phrase.is_empty() {
            return Err(HdWalletError::Bip39Error(Bip39Error::BadWordCount(0)).into());
        }
        PrivateKey::from_hd_wallet_path(path, phrase, passphrase)
    }
}

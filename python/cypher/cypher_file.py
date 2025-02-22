import hashlib
import base64
import pathlib
import argparse

from cryptography.fernet import Fernet


def make_key(secret: str) -> bytes:
    key = hashlib.sha256(secret.encode()).digest()
    return base64.urlsafe_b64encode(key)

def encrypt_file(path: pathlib.Path, secret: str) -> None:
    key = make_key(secret)
    cipher = Fernet(key)
    data = path.read_bytes()
    encrypted_file = path.parent / f"{path.name}.enc"
    encrypted_file.write_bytes(cipher.encrypt(data))
    print(f"Файл '{path}' успешно зашифрован.")

def decrypt_file(path: pathlib.Path, secret: str) -> None:
    assert path.suffix == ".enc"
    key = make_key(secret)
    cipher = Fernet(key)
    data = path.read_bytes()
    decrypted_data = cipher.decrypt(data)
    decrypted_file = pathlib.Path(path.stem)
    decrypted_file.write_bytes(decrypted_data)
    print(f"Файл '{path}' успешно расшифрован.")

def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description="Шифрование и дешифрование файла с использованием секретной фразы.")
    parser.add_argument('action', choices=['encrypt', 'decrypt'], help="Действие: шифрование или дешифрование.")
    parser.add_argument('file', help="Путь к файлу для шифрования или дешифрования.")
    parser.add_argument('--phrase', required=True, help="Секретная фраза для шифрования/дешифрования.")
    return parser


if __name__ == '__main__':
    parser = build_parser()
    args = parser.parse_args()

    path = pathlib.Path(args.file).resolve()
    match args.action:
        case 'encrypt':
            encrypt_file(path, args.phrase)
        case 'decrypt':
            decrypt_file(path, args.phrase)

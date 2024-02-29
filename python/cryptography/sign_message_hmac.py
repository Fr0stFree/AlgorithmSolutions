import hmac
import hashlib


def send_message() -> tuple[bytes, bytes]:
    message = 'Sending some bullshit message to the server.'
    decoded_message = message.encode(encoding)
    hash = hmac.new(shared_key, digestmod=algorithm)
    hash.update(decoded_message)
    return decoded_message, hash.digest()


def receive_message(message: bytes, hash: bytes) -> str:
    expected_hash = hmac.new(shared_key, digestmod=algorithm)
    expected_hash.update(message)
    is_valid = hmac.compare_digest(expected_hash.digest(), hash)
    if not is_valid:
        raise ValueError('Invalid message')

    encoded_message = message.decode(encoding)
    return encoded_message


if __name__ == "__main__":
    # Alice and Bob share a secret key. The shared key is a secret key that only Alice and Bob know.
    # The shared key is used to sign and verify messages.
    shared_key = b'secret'
    algorithm = hashlib.sha256
    encoding = 'utf-8'

    # Alice wants to send a message to Bob. She uses a shared key to sign the message.
    # The shared key is a secret key that only Alice and Bob know.
    # Alice uses the HMAC algorithm to sign the message.
    message, hash = send_message()

    # Bob receives the message and the hash. He uses the shared key to verify the message.
    result = receive_message(message, hash)

    print(f"Message: {result}")



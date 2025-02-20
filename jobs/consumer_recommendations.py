import json
import redis
import os

from generate_recommendations import generate_and_store_recommendations


redis_host = os.getenv('REDIS_HOST', 'localhost')
redis_port = int(os.getenv('REDIS_PORT', 6379))

redis_client = redis.Redis(host=redis_host, port=redis_port, decode_responses=True)
stream_name = 'stream:user:search_history'


def consume_messages():
    try:
        redis_client.xgroup_create(stream_name, 'group1', '$', mkstream=True)
    except redis.exceptions.ResponseError as e:
        pass

    while True:
        response = redis_client.xreadgroup('group1', 'consumer1', {stream_name: '>'}, block=0, count=1)
        
        for stream, messages in response:
            for message_id, message_data in messages:
                print(f'Consumer Group 1 (Consumer 1) processing: {message_data}')

                try:
                    user_id = message_data['user_id']
                    movie_ids = json.loads(message_data['movie_ids'])
                    generate_and_store_recommendations(user_id, movie_ids)
                except json.JSONDecodeError:
                    print('error: bad message format; message:', message_data)
                except Exception:
                    pass
                
                redis_client.xack(stream_name, 'group1', message_id)

if __name__ == '__main__':
    consume_messages()

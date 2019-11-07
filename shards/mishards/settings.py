import sys
import os

from environs import Env
env = Env()

FROM_EXAMPLE = env.bool('FROM_EXAMPLE', False)
if FROM_EXAMPLE:
    from dotenv import load_dotenv
    load_dotenv('./mishards/.env.example')
else:
    env.read_env()


DEBUG = env.bool('DEBUG', False)
MAX_RETRY = env.int('MAX_RETRY', 3)

LOG_LEVEL = env.str('LOG_LEVEL', 'DEBUG' if DEBUG else 'INFO')
LOG_PATH = env.str('LOG_PATH', '/tmp/mishards')
LOG_NAME = env.str('LOG_NAME', 'logfile')
TIMEZONE = env.str('TIMEZONE', 'UTC')

from utils.logger_helper import config
config(LOG_LEVEL, LOG_PATH, LOG_NAME, TIMEZONE)

SERVER_PORT = env.int('SERVER_PORT', 19530)
SERVER_TEST_PORT = env.int('SERVER_TEST_PORT', 19530)
WOSERVER = env.str('WOSERVER')


class TracingConfig:
    TRACING_SERVICE_NAME = env.str('TRACING_SERVICE_NAME', 'mishards')
    TRACING_VALIDATE = env.bool('TRACING_VALIDATE', True)
    TRACING_LOG_PAYLOAD = env.bool('TRACING_LOG_PAYLOAD', False)
    TRACING_CONFIG = {
        'sampler': {
            'type': env.str('TRACING_SAMPLER_TYPE', 'const'),
            'param': env.str('TRACING_SAMPLER_PARAM', "1"),
        },
        'local_agent': {
            'reporting_host': env.str('TRACING_REPORTING_HOST', '127.0.0.1'),
            'reporting_port': env.str('TRACING_REPORTING_PORT', '5775')
        },
        'logging': env.bool('TRACING_LOGGING', True)
    }
    DEFAULT_TRACING_CONFIG = {
        'sampler': {
            'type': env.str('TRACING_SAMPLER_TYPE', 'const'),
            'param': env.str('TRACING_SAMPLER_PARAM', "0"),
        }
    }


class DefaultConfig:
    SQLALCHEMY_DATABASE_URI = env.str('SQLALCHEMY_DATABASE_URI')
    SQL_ECHO = env.bool('SQL_ECHO', False)
    TRACER_PLUGIN_PATH = env.str('TRACER_PLUGIN_PATH', '')
    TRACER_CLASS_NAME = env.str('TRACER_CLASS_NAME', '')
    ROUTER_PLUGIN_PATH = env.str('ROUTER_PLUGIN_PATH', '')
    ROUTER_CLASS_NAME = env.str('ROUTER_CLASS_NAME', 'FileBasedHashRingRouter')
    DISCOVERY_PLUGIN_PATH = env.str('DISCOVERY_PLUGIN_PATH', '')
    DISCOVERY_CLASS_NAME = env.str('DISCOVERY_CLASS_NAME', 'static')


class TestingConfig(DefaultConfig):
    SQLALCHEMY_DATABASE_URI = env.str('SQLALCHEMY_DATABASE_TEST_URI', '')
    SQL_ECHO = env.bool('SQL_TEST_ECHO', False)
    TRACER_CLASS_NAME = env.str('TRACER_CLASS_TEST_NAME', '')
    ROUTER_CLASS_NAME = env.str('ROUTER_CLASS_TEST_NAME', 'FileBasedHashRingRouter')

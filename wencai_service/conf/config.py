import logging
import yaml

def init():
    with open('conf/config.yaml', 'r', encoding='utf-8') as f:
        return yaml.safe_load(f)
    
CONF=init()
logging.basicConfig(level=logging.DEBUG)
log = logging.getLogger(__name__)
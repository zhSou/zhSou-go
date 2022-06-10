import json

path_prefix = r'C:/Users/zzq/Desktop/新建文件夹'


def get_config():
    path_length = 256
    return {
        'pathLength': path_length,
        'dataFilePaths': [f'{path_prefix}/data/wukong_100m_{i}.dat' for i in range(path_length)],
        'dataIndexPaths': [f'{path_prefix}/index/wukong_100m_{i}.idx' for i in range(path_length)],
        'csvPaths': [f'{path_prefix}/input/wukong_100m_{i}.csv' for i in range(path_length)],
        'invertedIndexFilePath': f'{path_prefix}/inverted_index.inv',
        'dictPath': f'{path_prefix}/dict.dic',
        'stopWordPath': f'{path_prefix}/stop_words.txt',
        'importCsvCoroutines': 4,
        'makeInvertedIndexCoroutines': 8,
        'searchLruMaxCapacity': 20,
        'showMenu': True,
    }


def get_light_config():
    path_length = 1
    return {
        'pathLength': path_length,
        'dataFilePaths': [f'{path_prefix}/light/data/wukong_100m_{i}.dat' for i in range(path_length)],
        'dataIndexPaths': [f'{path_prefix}/light/index/wukong_100m_{i}.idx' for i in range(path_length)],
        'csvPaths': [f'{path_prefix}/input/wukong_100m_{i}.csv' for i in range(path_length)],
        'invertedIndexFilePath': f'{path_prefix}/light/inverted_index.inv',
        'dictPath': f'{path_prefix}/light/dict.dic',
        'stopWordPath': f'{path_prefix}/stop_words.txt',
        'importCsvCoroutines': 4,
        'makeInvertedIndexCoroutines': 8,
        'searchLruMaxCapacity': 20,
        'showMenu': False,
    }


with open('config.json', encoding='utf-8', mode='w') as f:
    json.dump(get_light_config(), f)

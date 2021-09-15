#!/usr/bin/env python

import yaml


class Nested:
    yaml_tag = u'!Nested'

    def __init__(self, id, value):
        self.id = id
        self.value = value

    @classmethod
    def to_yaml(cls, dumper: yaml.SafeDumper, data) -> yaml.nodes.MappingNode:
        mapping = {
            "id": data.id,
            "val": data.value,
        }
        return dumper.represent_mapping(cls.yaml_tag, mapping)


class Example:
    yaml_tag = u'!Example'

    def __init__(self):
        id = "123-456-abc"
        self._nested = Nested(id, 9876)
        self._number = 123
        self._id = id

    @property
    def nested(self):
        return self.nested

    @property
    def number(self):
        return self._number

    @property
    def id(self):
        return self._id

    @classmethod
    def to_yaml(cls, dumper: yaml.SafeDumper, data) -> yaml.nodes.MappingNode:
        mapping = {
            "id": data.id,
            "number": data.number,
            "nested": data._nested,
        }
        return dumper.represent_mapping(cls.yaml_tag, mapping)


def Example_representer(dumper, data):
    serializedData = data.id + " | " + str(data.number)
    return dumper.represent_scalar('!Example', serializedData)


def Example2_representer(dumper: yaml.SafeDumper, obj: Example) -> yaml.nodes.MappingNode:
    return dumper.represent_mapping('!Example', {
        "id": obj.id,
        "number": obj.number,
        "field1": obj.field1,
    })


def noop(self, *args, **kw):
    pass


# yaml.add_representer(Example, Example2_representer)
yaml.SafeDumper.add_multi_representer(Example, Example.to_yaml)
yaml.SafeDumper.add_multi_representer(Nested, Nested.to_yaml)
yaml.emitter.Emitter.process_tag = noop


if __name__ == "__main__":

    obj = Example()

    ystr = yaml.safe_dump(obj,
                          explicit_start=True,
                          explicit_end=True)

    print(ystr)

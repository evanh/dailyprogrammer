import json
import urllib2

HOSTNAME = 'https://challenge.opencare.com'
NUCLEOTIDES = ['A', 'C', 'G', 'T']

def get_project_info():
    response = urllib2.urlopen(HOSTNAME + "/start")
    return json.loads(response.read())

def build_tree(project_info):
    tree = []
    for i in xrange(project_info['resultLength']):
        at_location = []
        for n in NUCLEOTIDES:
            queryparams = "?location=%d&uniq=%s&nucleotide=%s" % (i, project_info['uniq'], n)
            response = urllib2.urlopen(HOSTNAME + "/location" + queryparams)
            ret = json.loads(response.read())
            if ret['expressed']:
                at_location.append(n)
        tree.append(at_location)

    return tree

def get_genome(genome_so_far, tree, project_info):
    if len(tree) == 1:
        matching = []
        for n in tree[0]:
            if project_info['partialSequence'] in (genome_so_far + n):
                matching.append(genome_so_far + n)
        return matching

    matching = []
    for n in tree[0]:
        matching.extend(get_genome(genome_so_far + n, tree[1:], project_info))
    return matching

if __name__ == "__main__":
    project_info = get_project_info()
    print project_info['uniq']
    genome_tree = build_tree(project_info)
    matches = get_genome("", genome_tree, project_info)
    print json.dumps(matches)
    data = {
        "sequences": matches,
        "uniq": project_info['uniq'],
    }
    req = urllib2.Request(HOSTNAME + "/check", json.dumps(data), {"Content-Type": "application/json"})
    
    try:
        response = urllib2.urlopen(req)
        if response.code != 200:
            print response.read()
    except Exception as e:
        print "%s, %s was not correct" % (e, data)


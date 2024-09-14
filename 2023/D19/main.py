import re


def parse_workflow(workflow_line: str) -> tuple[str, list[dict]]:
    match = re.match(r'(\w+)\{(.*)\}', workflow_line)

    workflow_name = match.group(1)
    operations = []
    for operation in match.group(2).split(','):
        tmp = {}
        if '>' in operation:
            [evaluation, destination] = operation.split(':')
            [prop_name, amount] = evaluation.split('>')
            operations.append({'amount': int(amount), 'prop_name': prop_name, 'operator': '>', 'destination': destination})
        elif '<' in operation:
            [evaluation, destination] = operation.split(':')
            [prop_name, amount] = evaluation.split('<')
            operations.append({'amount': int(amount), 'prop_name': prop_name, 'operator': '<', 'destination': destination})
        else:
            operations.append({'operator': None, 'destination': operation})

    return workflow_name, operations


if __name__ == '__main__':
    # input_file = open('./test.txt')
    input_file = open('./input.txt')

    [workflows_input, parts_input] = input_file.read().strip().split('\n\n')
    input_file.close()

    workflows = {}
    processed_parts = {
        'A': [],
        'R': [],
    }
    for workflow in workflows_input.split('\n'):
        [workflow_name, operations] = parse_workflow(workflow)
        workflows[workflow_name] = operations

    parts = []
    for part_input in parts_input.split('\n'):
        part = {}
        for prop in part_input.replace('{', '').replace('}', '').split(','):
            [name, value] = prop.split('=')
            part[name] = int(value)

        parts.append(part)

    print(workflows)

    parts_to_process = len(parts)
    for part in parts:
        processed = False
        workflow_name = 'in'
        path = []
        while not processed:
            path.append(workflow_name)
            if workflow_name in ['A', 'R']:
                processed_parts[workflow_name].append(part)
                processed = True
            else:
                operations = workflows[workflow_name]
                for operation in operations:
                    if operation.get('operator') == '>':
                        if part[operation.get('prop_name')] > operation.get('amount'):
                            workflow_name = operation.get('destination')
                            break
                    elif operation.get('operator') == '<':
                        if part[operation.get('prop_name')] < operation.get('amount'):
                            workflow_name = operation.get('destination')
                            break
                    elif operation.get('operator') is None:
                        workflow_name = operation.get('destination')
                        break
        print(f"{part}: {'->'.join(path)}")

    rating = 0
    for processed_part in processed_parts['A']:
        for part_rating in processed_part:
            rating += processed_part[part_rating]

    print(rating)

# m > 838, m > 2090, m > 1548
# a < 1716, a <= 3333,
# x <= 2440, x < 1416, x > 2662
# s > 3448,

test = [1,1,1,2,3]

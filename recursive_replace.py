def recursive_replace(data: dict, target: str, replacement: str) -> None:
  """Recursively walks a dict structure searching for a target string value, performs the replace in place

  :param data: The dict to walk
  :param target: The target to scan for
  :param replacement: The required value to replace target with
  :return: None
  """
  if isinstance(data, dict):
    for key, value in data.items():
      if isinstance(value, str):
        data[key] = re.sub(target, replacement, value)
      elif isinstance(value, dict):
        recursive_replace(value, target, replacement)
      elif isinstance(value, list):
        recursive_replace(value, target, replacement)
  elif isinstance(data, list):
    for i in range(len(data)):
      if isinstance(data[i], str):
        data[i] = re.sub(target, replacement, data[i])
      else:
        recursive_replace(data[i], target, replacement)

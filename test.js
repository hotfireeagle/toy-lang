const asyncPool = async (limit, datas, iteratorFun) => {
  const running_task = []
  const all_task = []

  for (const data of datas) {
    const promise1 = iteratorFun(data)
    all_task.push(promise1)

    if (limit < datas.length) {
      const promise2 = promise1.then(() => {
        const idx = running_task.indexOf(promise2)
        running_task.splice(idx, 1)
      })
      running_task.push(promise2)
    }

    if (limit == running_task.length) {
      await Promise.race(running_task)
    }
  }

  await Promise.all(all_task)
}
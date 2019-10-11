function main(){
  const cycles = 1000000000 // 1,000,000,000
  let start = Date.now()
  
  for(let i = 0; i < cycles; i++){
    // Empty loop
  }
  
  let end = Date.now()
  let duration = (end - start) / 1000

  console.info(`Javascript looped ${cycles} in ${duration} seconds`)
}

main();

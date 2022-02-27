 - A struct to hold information about the target and the CAS protein being used.
 - A substruct to hold information about the CAS protein (PAM, references, functionality (cut, tag))
 - Maybe a small embedded DB of common CAS variants and their properties.
 - A basic heuristic for determining the viability of a potential gRNA candidate. This should include:
 -- Checking for off-target hits in the target organism ("this target is a bad idea because the DNA exists in different places" vs "this gRNA is a bad idea ")
 	> Might want to have both "this target is invalid" and "this gRNA is invalid"
 -- Checking that the gRNA itself can be synthesized
 -- Checking that gRNA GC content is within the 40-80% range.
 -- sgRNA should be between 17-24 nucleotides



Turns out the NCBI API is mildly horrible, and needs to be iteratively re-polled to get a valid response. Specifically, it (1) parses the horrible HTML loading page, and then (2) iteratively makes new requests until it gets a content-ful one, and then (3) returns that

I was able to get a valid XML response using the logic defined in biopython here: 

https://github.com/biopython/biopython/blob/4d37d8165a9b7c60cc298e9b67a1bfe57c825840/Bio/Blast/NCBIWWW.py#L276


```
from Bio.Blast import NCBIWWW
result = NCBIWWW.qblast(program='blastn', database='nt', 	sequence='TGATCCACTATTAAACCGGCTTTAGGTTGtttttttGTTCTGTCATTTTGCTTTTTACC', format_type='XML')
print(result.read())
```

It might be worth pulling this functionality from Python directly rather than trying to re-implement it in Go? 



Spreadsheet with CAS info: https://docs.google.com/spreadsheets/d/1sl7xRxUueIf_c77jnliS7tTwLaM83zVcV30Xj8y_36c/edit?usp=sharing

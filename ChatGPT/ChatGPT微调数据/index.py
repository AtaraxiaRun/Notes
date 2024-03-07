import openai

openai.api_key = "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"

model_info = openai.FineTune.retrieve(id="ft-gm3VC33qcsOybFgV0Ra8ZXFu")
print(model_info)

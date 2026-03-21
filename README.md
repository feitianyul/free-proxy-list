**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-21 06:35:34 UTC（2026-03-21 14:35:34 UTC+8）

**代理总数：267**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 267 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 267 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 658ms | 否 | ✓ 848ms | ✓ 938ms | ✓ 971ms | http |
| 137.220.150.170:6005 | ✓ 1187ms | 否 | ✓ 990ms | ✓ 1144ms | ✓ 953ms | http |
| 219.117.204.211:7799 | ✓ 817ms | 否 | 否 | ✓ 1572ms | ✓ 953ms | http |
| 147.161.239.240:8800 | ✓ 1371ms | ✓ 1816ms | ✓ 1252ms | ✓ 1743ms | ✓ 1685ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1850ms | 否 | ✓ 1275ms | ✓ 1078ms | http |
| 35.225.22.61:80 | ✓ 323ms | 否 | 否 | ✓ 1191ms | ✓ 1194ms | http |
| 174.138.24.77:1080 | ✓ 1067ms | 否 | ✓ 1067ms | ✓ 1219ms | ✓ 996ms | http |
| 137.220.150.22:6005 | ✓ 708ms | 否 | ✓ 951ms | ✓ 1254ms | ✓ 1107ms | http |
| 149.28.157.177:1080 | ✓ 754ms | 否 | ✓ 1931ms | 否 | ✓ 1340ms | http |
| 163.44.126.97:3128 | ✓ 1472ms | 否 | 否 | ✓ 1441ms | ✓ 1712ms | http |
| 115.231.181.40:8128 | ✓ 951ms | 否 | ✓ 898ms | ✓ 1838ms | 否 | http |
| 133.242.138.34:8100 | 否 | 否 | ✓ 648ms | ✓ 1965ms | ✓ 988ms | http |
| 94.16.114.3:40000 | ✓ 1292ms | 否 | ✓ 1487ms | 否 | ✓ 1924ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1312ms | ✓ 1466ms | 否 | ✓ 923ms | http |
| 167.103.31.122:8800 | ✓ 1677ms | 否 | ✓ 1878ms | 否 | ✓ 1849ms | http |
| 38.34.179.229:8445 | ✓ 500ms | ✓ 1570ms | ✓ 219ms | ✓ 857ms | ✓ 966ms | http |
| 38.34.179.75:8453 | ✓ 293ms | ✓ 1072ms | ✓ 505ms | ✓ 1349ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1943ms | 否 | ✓ 1172ms | ✓ 1356ms | ✓ 1226ms | http |
| 45.88.0.111:3128 | ✓ 1108ms | 否 | ✓ 1850ms | 否 | ✓ 1205ms | http |
| 38.145.208.189:8446 | 否 | ✓ 1591ms | ✓ 632ms | ✓ 795ms | ✓ 512ms | http |
| 137.220.151.110:6005 | ✓ 825ms | ✓ 1727ms | ✓ 969ms | ✓ 1145ms | ✓ 1478ms | http |
| 91.238.105.64:2024 | ✓ 876ms | ✓ 1768ms | ✓ 1109ms | ✓ 1742ms | ✓ 1362ms | http |
| 120.92.212.16:8890 | ✓ 1146ms | 否 | ✓ 1000ms | 否 | ✓ 944ms | http |
| 38.145.208.244:8448 | ✓ 156ms | ✓ 1575ms | ✓ 152ms | ✓ 725ms | ✓ 549ms | http |
| 164.90.155.209:3128 | ✓ 367ms | ✓ 604ms | ✓ 772ms | ✓ 667ms | 否 | http |
| 45.149.92.147:5001 | ✓ 633ms | 否 | ✓ 593ms | ✓ 761ms | ✓ 611ms | http |
| 137.220.150.104:6005 | ✓ 709ms | 否 | ✓ 1007ms | ✓ 1140ms | ✓ 1262ms | http |
| 160.250.4.245:1 | 否 | 否 | ✓ 1157ms | ✓ 1141ms | ✓ 936ms | http |
| 43.225.185.4:8000 | ✓ 909ms | 否 | ✓ 1522ms | ✓ 1286ms | ✓ 1277ms | http |
| 106.75.15.167:7890 | ✓ 1979ms | ✓ 1858ms | 否 | 否 | ✓ 1409ms | http |
| 59.46.216.131:30001 | ✓ 974ms | ✓ 1574ms | 否 | ✓ 1385ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1792ms | 否 | ✓ 1687ms | ✓ 1169ms | 否 | http |
| 38.145.208.167:8445 | 否 | 否 | ✓ 199ms | ✓ 725ms | ✓ 521ms | http |
| 38.34.179.85:8444 | 否 | ✓ 706ms | ✓ 113ms | ✓ 727ms | ✓ 639ms | http |
| 49.156.44.114:8080 | 否 | 否 | ✓ 1515ms | ✓ 1324ms | ✓ 1529ms | http |
| 38.34.183.16:8452 | ✓ 208ms | ✓ 681ms | ✓ 94ms | ✓ 668ms | ✓ 545ms | http |
| 38.145.220.8:8452 | ✓ 194ms | ✓ 673ms | ✓ 100ms | ✓ 689ms | ✓ 554ms | http |
| 38.34.179.89:8444 | ✓ 200ms | ✓ 739ms | ✓ 98ms | ✓ 673ms | ✓ 517ms | http |
| 45.136.131.64:8444 | ✓ 412ms | ✓ 986ms | ✓ 88ms | ✓ 674ms | ✓ 515ms | http |
| 45.136.131.67:8444 | ✓ 400ms | ✓ 1411ms | ✓ 100ms | ✓ 676ms | ✓ 519ms | http |
| 38.34.183.221:8452 | ✓ 194ms | ✓ 1430ms | ✓ 138ms | ✓ 685ms | ✓ 538ms | http |
| 38.34.179.91:8444 | ✓ 195ms | 否 | ✓ 103ms | ✓ 783ms | ✓ 531ms | http |
| 45.136.131.68:8444 | ✓ 400ms | 否 | ✓ 116ms | ✓ 707ms | ✓ 565ms | http |
| 38.145.220.33:8448 | ✓ 455ms | ✓ 809ms | ✓ 532ms | ✓ 1121ms | ✓ 534ms | http |
| 148.153.56.51:80 | 否 | ✓ 1596ms | ✓ 776ms | ✓ 1850ms | ✓ 673ms | http |
| 38.34.179.99:8448 | ✓ 1256ms | 否 | ✓ 253ms | ✓ 725ms | ✓ 879ms | http |
| 38.145.218.235:8445 | ✓ 1271ms | ✓ 1961ms | ✓ 275ms | ✓ 749ms | 否 | http |
| 38.34.179.190:8450 | ✓ 959ms | ✓ 1054ms | ✓ 675ms | ✓ 817ms | ✓ 1010ms | http |
| 45.88.0.114:3128 | ✓ 1050ms | ✓ 1941ms | ✓ 1059ms | ✓ 1972ms | 否 | http |
| 153.126.195.219:5080 | ✓ 653ms | 否 | ✓ 1169ms | ✓ 1466ms | ✓ 823ms | http |
| 137.220.150.152:6005 | ✓ 967ms | 否 | ✓ 733ms | ✓ 1472ms | 否 | http |
| 45.88.0.113:3128 | ✓ 1051ms | 否 | ✓ 1137ms | ✓ 1518ms | 否 | http |
| 45.88.0.116:3128 | ✓ 1248ms | 否 | ✓ 1577ms | 否 | ✓ 1175ms | http |
| 45.88.0.115:3128 | ✓ 1674ms | 否 | ✓ 1654ms | ✓ 1514ms | ✓ 1151ms | http |
| 45.88.0.117:3128 | ✓ 1677ms | ✓ 1600ms | ✓ 1636ms | 否 | 否 | http |
| 103.113.70.189:1081 | ✓ 1353ms | ✓ 1241ms | ✓ 1213ms | ✓ 1273ms | ✓ 949ms | http |
| 45.136.130.182:8450 | ✓ 723ms | ✓ 1215ms | ✓ 812ms | ✓ 1719ms | ✓ 749ms | http |
| 38.145.218.229:8450 | ✓ 89ms | ✓ 1140ms | ✓ 98ms | ✓ 765ms | ✓ 1524ms | http |
| 38.34.179.33:8445 | ✓ 727ms | ✓ 1325ms | ✓ 184ms | ✓ 944ms | ✓ 562ms | http |
| 38.34.179.98:8453 | ✓ 140ms | ✓ 835ms | ✓ 97ms | ✓ 880ms | ✓ 765ms | http |
| 45.136.131.30:8451 | ✓ 107ms | ✓ 1668ms | ✓ 127ms | ✓ 696ms | ✓ 544ms | http |
| 45.136.131.32:8451 | ✓ 295ms | ✓ 764ms | ✓ 91ms | ✓ 1222ms | ✓ 870ms | http |
| 45.136.131.35:8452 | ✓ 236ms | ✓ 1600ms | ✓ 152ms | 否 | ✓ 515ms | http |
| 45.136.131.34:8451 | ✓ 146ms | ✓ 1651ms | ✓ 141ms | ✓ 763ms | ✓ 889ms | http |
| 38.34.179.72:8445 | ✓ 141ms | ✓ 731ms | 否 | ✓ 917ms | ✓ 691ms | http |
| 45.136.131.25:8450 | ✓ 1073ms | ✓ 1488ms | ✓ 475ms | ✓ 1638ms | ✓ 1346ms | http |
| 45.136.130.171:8445 | ✓ 461ms | ✓ 1924ms | ✓ 83ms | ✓ 705ms | ✓ 636ms | http |
| 38.34.183.164:8453 | ✓ 343ms | ✓ 806ms | ✓ 176ms | ✓ 785ms | 否 | http |
| 38.34.179.80:8452 | ✓ 423ms | ✓ 1758ms | ✓ 228ms | ✓ 1843ms | ✓ 667ms | http |
| 38.34.179.94:8451 | ✓ 197ms | ✓ 1458ms | ✓ 1143ms | 否 | ✓ 1114ms | http |
| 38.34.179.86:8452 | ✓ 192ms | ✓ 850ms | ✓ 101ms | ✓ 683ms | ✓ 526ms | http |
| 45.136.130.177:8448 | ✓ 115ms | ✓ 939ms | ✓ 95ms | ✓ 715ms | ✓ 508ms | http |
| 38.34.179.194:8446 | ✓ 149ms | ✓ 631ms | ✓ 100ms | ✓ 697ms | ✓ 600ms | http |
| 45.136.130.178:8449 | ✓ 82ms | ✓ 865ms | ✓ 86ms | ✓ 676ms | ✓ 492ms | http |
| 38.34.179.155:8452 | ✓ 141ms | ✓ 782ms | ✓ 169ms | ✓ 704ms | ✓ 648ms | http |
| 38.34.179.186:8444 | ✓ 250ms | ✓ 626ms | ✓ 125ms | ✓ 929ms | ✓ 521ms | http |
| 45.136.130.188:8449 | ✓ 92ms | ✓ 1166ms | ✓ 102ms | ✓ 705ms | ✓ 581ms | http |
| 45.136.130.195:8444 | ✓ 546ms | ✓ 612ms | ✓ 107ms | ✓ 686ms | ✓ 737ms | http |
| 38.34.179.165:8446 | ✓ 116ms | ✓ 681ms | ✓ 294ms | ✓ 678ms | ✓ 994ms | http |
| 38.34.178.7:8452 | ✓ 148ms | ✓ 1486ms | ✓ 94ms | ✓ 684ms | ✓ 519ms | http |
| 38.34.179.173:8452 | ✓ 142ms | ✓ 635ms | ✓ 502ms | ✓ 687ms | ✓ 840ms | http |
| 38.34.179.87:8451 | ✓ 646ms | ✓ 982ms | ✓ 92ms | ✓ 707ms | ✓ 518ms | http |
| 38.34.179.104:8446 | ✓ 125ms | ✓ 1346ms | ✓ 114ms | ✓ 735ms | ✓ 561ms | http |
| 38.34.178.141:8453 | ✓ 157ms | ✓ 1279ms | ✓ 149ms | ✓ 718ms | ✓ 531ms | http |
| 38.34.179.101:8446 | ✓ 123ms | ✓ 1167ms | ✓ 314ms | ✓ 748ms | ✓ 577ms | http |
| 38.145.208.185:8449 | ✓ 138ms | ✓ 713ms | ✓ 111ms | ✓ 703ms | ✓ 806ms | http |
| 38.34.179.106:8446 | ✓ 363ms | ✓ 1161ms | ✓ 107ms | ✓ 696ms | ✓ 580ms | http |
| 45.136.130.193:8444 | ✓ 82ms | ✓ 1212ms | ✓ 90ms | ✓ 696ms | ✓ 1007ms | http |
| 38.34.178.186:8451 | ✓ 146ms | ✓ 690ms | ✓ 351ms | ✓ 757ms | ✓ 763ms | http |
| 38.34.179.6:8449 | ✓ 148ms | ✓ 1503ms | ✓ 93ms | ✓ 671ms | ✓ 551ms | http |
| 38.34.179.92:8451 | ✓ 418ms | ✓ 1785ms | ✓ 102ms | ✓ 695ms | ✓ 536ms | http |
| 38.34.179.88:8446 | ✓ 177ms | ✓ 1553ms | ✓ 125ms | ✓ 897ms | ✓ 509ms | http |
| 38.34.178.241:8453 | ✓ 132ms | ✓ 851ms | ✓ 92ms | ✓ 1762ms | ✓ 527ms | http |
| 38.34.179.39:8452 | ✓ 333ms | ✓ 1159ms | ✓ 143ms | ✓ 729ms | ✓ 492ms | http |
| 45.136.131.62:8449 | ✓ 218ms | ✓ 1235ms | ✓ 740ms | ✓ 659ms | ✓ 721ms | http |
| 38.34.178.193:8452 | ✓ 151ms | ✓ 700ms | ✓ 309ms | ✓ 695ms | ✓ 538ms | http |
| 38.34.179.204:8446 | ✓ 176ms | ✓ 686ms | ✓ 1355ms | ✓ 709ms | ✓ 1018ms | http |
| 38.34.179.40:8446 | ✓ 214ms | ✓ 1747ms | ✓ 113ms | ✓ 843ms | ✓ 544ms | http |
| 38.34.179.25:8444 | ✓ 103ms | ✓ 1915ms | ✓ 117ms | ✓ 692ms | ✓ 542ms | http |
| 38.34.179.23:8444 | ✓ 241ms | ✓ 720ms | ✓ 97ms | ✓ 745ms | ✓ 570ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)

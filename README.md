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

最后更新：2026-03-21 11:24:41 UTC（2026-03-21 19:24:41 UTC+8）

**代理总数：253**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 252 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 253 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1670ms | 否 | ✓ 975ms | ✓ 1161ms | ✓ 991ms | http |
| 167.103.34.108:8800 | ✓ 1624ms | 否 | ✓ 1594ms | ✓ 1554ms | ✓ 1583ms | http |
| 101.47.73.135:3128 | ✓ 1869ms | 否 | ✓ 1150ms | ✓ 1487ms | ✓ 1389ms | http |
| 106.75.15.167:7890 | ✓ 1249ms | ✓ 1399ms | ✓ 1008ms | ✓ 1369ms | 否 | http |
| 45.167.124.52:8080 | ✓ 1089ms | 否 | ✓ 1680ms | 否 | ✓ 1607ms | http |
| 137.220.150.22:6005 | ✓ 1765ms | 否 | ✓ 860ms | ✓ 1202ms | ✓ 1037ms | http |
| 193.23.200.251:10808 | ✓ 1376ms | 否 | ✓ 1019ms | 否 | ✓ 1333ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1419ms | ✓ 1397ms | ✓ 1134ms | http |
| 167.103.31.122:8800 | ✓ 1837ms | 否 | ✓ 1892ms | ✓ 1870ms | 否 | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 243ms | ✓ 913ms | ✓ 785ms | http |
| 120.92.212.16:8890 | ✓ 1393ms | ✓ 1634ms | ✓ 1124ms | ✓ 1635ms | ✓ 1312ms | http |
| 37.187.109.70:10111 | ✓ 1451ms | 否 | ✓ 1315ms | ✓ 1697ms | ✓ 1524ms | http |
| 142.171.224.229:7890 | ✓ 1129ms | 否 | ✓ 907ms | ✓ 1430ms | 否 | http |
| 147.161.239.240:8800 | ✓ 1276ms | 否 | ✓ 782ms | ✓ 1955ms | ✓ 1324ms | http |
| 101.43.127.100:8877 | ✓ 1093ms | 否 | ✓ 1095ms | ✓ 1989ms | ✓ 958ms | http |
| 217.174.244.117:3129 | ✓ 1023ms | ✓ 1686ms | ✓ 1455ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1065ms | 否 | 否 | ✓ 1346ms | ✓ 1352ms | http |
| 172.212.68.37:3128 | 否 | 否 | ✓ 1444ms | ✓ 1240ms | ✓ 1260ms | http |
| 192.71.213.85:5678 | ✓ 1294ms | 否 | ✓ 1634ms | ✓ 1756ms | 否 | http |
| 137.220.150.170:6005 | ✓ 1479ms | 否 | ✓ 1334ms | ✓ 1418ms | ✓ 1004ms | http |
| 137.220.151.110:6005 | ✓ 1631ms | 否 | ✓ 1152ms | ✓ 1903ms | ✓ 1705ms | http |
| 137.220.150.152:6005 | ✓ 1113ms | 否 | ✓ 1172ms | ✓ 1475ms | ✓ 1028ms | http |
| 47.77.193.180:1080 | ✓ 245ms | ✓ 1543ms | ✓ 351ms | ✓ 829ms | ✓ 629ms | http |
| 144.31.79.117:8888 | ✓ 1586ms | 否 | ✓ 1205ms | 否 | ✓ 1991ms | http |
| 104.129.202.127:10810 | 否 | 否 | ✓ 751ms | ✓ 929ms | ✓ 848ms | http |
| 104.129.202.127:12354 | ✓ 591ms | ✓ 939ms | ✓ 556ms | ✓ 987ms | ✓ 710ms | http |
| 219.117.204.211:7799 | ✓ 1521ms | ✓ 1731ms | 否 | ✓ 1001ms | ✓ 777ms | http |
| 181.78.44.63:999 | ✓ 1423ms | ✓ 1521ms | ✓ 1210ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1602ms | 否 | 否 | ✓ 1631ms | ✓ 867ms | http |
| 59.46.216.131:30001 | ✓ 1063ms | 否 | ✓ 1232ms | 否 | ✓ 1160ms | http |
| 166.88.55.83:7890 | ✓ 774ms | ✓ 1215ms | ✓ 749ms | ✓ 951ms | ✓ 1382ms | http |
| 139.159.99.242:8080 | ✓ 917ms | ✓ 1126ms | ✓ 1000ms | 否 | 否 | http |
| 210.45.70.16:7895 | ✓ 1544ms | ✓ 1867ms | ✓ 1537ms | ✓ 1559ms | ✓ 1849ms | http |
| 202.141.161.53:30001 | 否 | ✓ 1432ms | ✓ 1458ms | 否 | ✓ 1469ms | http |
| 103.84.95.54:7890 | ✓ 1184ms | 否 | ✓ 860ms | 否 | ✓ 764ms | http |
| 137.220.150.104:6005 | ✓ 1155ms | 否 | ✓ 896ms | ✓ 1191ms | ✓ 945ms | http |
| 140.82.34.41:1080 | ✓ 946ms | 否 | ✓ 1086ms | ✓ 1871ms | ✓ 1127ms | http |
| 45.93.29.147:6005 | ✓ 907ms | ✓ 1693ms | ✓ 1737ms | ✓ 1216ms | 否 | http |
| 38.34.179.61:8445 | 否 | ✓ 861ms | ✓ 421ms | ✓ 889ms | ✓ 769ms | http |
| 38.34.179.104:8446 | 否 | ✓ 1592ms | ✓ 245ms | ✓ 898ms | ✓ 993ms | http |
| 38.34.179.101:8446 | 否 | 否 | ✓ 254ms | ✓ 859ms | ✓ 686ms | http |
| 45.186.6.104:3128 | ✓ 1063ms | ✓ 1891ms | ✓ 1857ms | 否 | 否 | http |
| 45.136.130.185:8444 | 否 | ✓ 1582ms | ✓ 420ms | ✓ 1066ms | ✓ 676ms | http |
| 38.145.208.172:8448 | 否 | 否 | ✓ 311ms | ✓ 893ms | ✓ 691ms | http |
| 91.238.105.64:2024 | ✓ 1097ms | 否 | ✓ 1862ms | ✓ 1957ms | ✓ 1549ms | http |
| 106.14.203.63:3333 | ✓ 975ms | ✓ 1694ms | 否 | 否 | ✓ 1736ms | http |
| 91.233.223.147:3128 | ✓ 1685ms | 否 | ✓ 1806ms | 否 | ✓ 1603ms | http |
| 38.34.179.98:8453 | ✓ 392ms | ✓ 1063ms | ✓ 233ms | ✓ 890ms | ✓ 774ms | http |
| 38.34.179.97:8448 | ✓ 401ms | ✓ 1004ms | ✓ 247ms | ✓ 902ms | ✓ 807ms | http |
| 38.34.179.102:8443 | ✓ 400ms | ✓ 1057ms | ✓ 229ms | ✓ 858ms | ✓ 821ms | http |
| 38.34.179.96:8451 | ✓ 397ms | ✓ 1478ms | ✓ 228ms | ✓ 867ms | ✓ 692ms | http |
| 38.34.179.103:8443 | ✓ 393ms | 否 | ✓ 248ms | ✓ 876ms | ✓ 711ms | http |
| 160.250.4.245:1 | ✓ 971ms | 否 | ✓ 1500ms | 否 | ✓ 1199ms | http |
| 133.242.138.34:8100 | ✓ 1750ms | ✓ 1999ms | 否 | 否 | ✓ 1506ms | http |
| 45.136.198.40:3128 | ✓ 1185ms | ✓ 1967ms | 否 | 否 | ✓ 1764ms | http |
| 103.39.51.190:8080 | ✓ 1954ms | 否 | 否 | ✓ 1552ms | ✓ 1514ms | http |
| 38.145.208.242:8451 | ✓ 252ms | ✓ 894ms | ✓ 250ms | ✓ 902ms | ✓ 671ms | http |
| 38.34.179.78:8445 | ✓ 324ms | ✓ 1351ms | ✓ 375ms | ✓ 923ms | ✓ 692ms | http |
| 38.34.179.57:8453 | ✓ 442ms | ✓ 858ms | ✓ 237ms | ✓ 853ms | ✓ 654ms | http |
| 38.34.179.54:8446 | ✓ 479ms | ✓ 1368ms | ✓ 263ms | ✓ 949ms | ✓ 667ms | http |
| 38.34.179.173:8452 | ✓ 345ms | ✓ 1786ms | ✓ 742ms | ✓ 865ms | ✓ 688ms | http |
| 38.34.178.186:8451 | ✓ 719ms | ✓ 1586ms | ✓ 1005ms | ✓ 884ms | ✓ 677ms | http |
| 45.136.131.62:8449 | ✓ 282ms | ✓ 1683ms | ✓ 898ms | ✓ 984ms | ✓ 797ms | http |
| 45.136.130.192:8451 | ✓ 424ms | 否 | ✓ 508ms | ✓ 1256ms | ✓ 861ms | http |
| 38.34.179.165:8446 | ✓ 762ms | ✓ 825ms | ✓ 271ms | ✓ 869ms | ✓ 718ms | http |
| 38.34.179.39:8452 | ✓ 785ms | ✓ 886ms | ✓ 251ms | ✓ 863ms | ✓ 754ms | http |
| 38.34.179.40:8446 | ✓ 791ms | ✓ 810ms | ✓ 252ms | ✓ 875ms | ✓ 990ms | http |
| 45.136.130.178:8449 | ✓ 769ms | ✓ 1743ms | ✓ 283ms | ✓ 878ms | ✓ 697ms | http |
| 38.34.179.51:8449 | ✓ 768ms | 否 | ✓ 309ms | ✓ 900ms | ✓ 758ms | http |
| 38.34.179.48:8449 | ✓ 764ms | 否 | ✓ 302ms | ✓ 909ms | ✓ 755ms | http |
| 24.144.86.173:1080 | ✓ 570ms | ✓ 1106ms | ✓ 1155ms | ✓ 1028ms | ✓ 839ms | http |
| 45.136.130.189:8451 | ✓ 1092ms | ✓ 956ms | ✓ 600ms | ✓ 1665ms | ✓ 1435ms | http |
| 45.136.130.184:8449 | ✓ 736ms | ✓ 1928ms | ✓ 235ms | ✓ 880ms | ✓ 697ms | http |
| 38.145.220.11:8445 | ✓ 473ms | ✓ 1574ms | ✓ 234ms | ✓ 881ms | ✓ 691ms | http |
| 45.136.130.177:8448 | ✓ 289ms | ✓ 873ms | ✓ 272ms | ✓ 920ms | ✓ 675ms | http |
| 45.136.130.193:8444 | ✓ 351ms | ✓ 910ms | ✓ 507ms | ✓ 918ms | ✓ 715ms | http |
| 38.34.179.178:8445 | ✓ 496ms | 否 | ✓ 271ms | ✓ 940ms | ✓ 755ms | http |
| 45.136.130.195:8444 | ✓ 385ms | ✓ 1665ms | ✓ 274ms | ✓ 876ms | ✓ 687ms | http |
| 45.136.130.186:8451 | ✓ 268ms | 否 | ✓ 253ms | ✓ 922ms | ✓ 734ms | http |
| 38.145.203.108:8445 | ✓ 881ms | ✓ 1696ms | ✓ 365ms | ✓ 1846ms | ✓ 1435ms | http |
| 38.34.179.49:8450 | ✓ 1220ms | ✓ 1282ms | ✓ 296ms | ✓ 1081ms | ✓ 1150ms | http |
| 38.34.179.203:8451 | 否 | 否 | ✓ 405ms | ✓ 1308ms | ✓ 782ms | http |
| 38.145.220.96:8443 | ✓ 571ms | ✓ 833ms | ✓ 449ms | ✓ 846ms | ✓ 672ms | http |
| 38.145.220.103:8443 | ✓ 573ms | ✓ 1243ms | ✓ 231ms | ✓ 867ms | ✓ 661ms | http |
| 38.145.203.87:8443 | ✓ 573ms | ✓ 1271ms | ✓ 233ms | ✓ 853ms | ✓ 647ms | http |
| 38.145.203.76:8443 | ✓ 575ms | ✓ 1244ms | ✓ 230ms | ✓ 861ms | ✓ 648ms | http |
| 38.145.208.181:8445 | ✓ 580ms | ✓ 1272ms | ✓ 269ms | ✓ 860ms | ✓ 710ms | http |
| 38.145.220.100:8443 | ✓ 574ms | 否 | ✓ 234ms | ✓ 873ms | ✓ 839ms | http |
| 38.34.179.25:8444 | ✓ 579ms | 否 | ✓ 269ms | ✓ 902ms | ✓ 782ms | http |
| 38.34.179.192:8450 | 否 | 否 | ✓ 1885ms | ✓ 967ms | ✓ 679ms | http |
| 38.34.179.190:8450 | 否 | 否 | ✓ 1885ms | ✓ 1031ms | ✓ 720ms | http |
| 20.78.213.56:80 | ✓ 1307ms | 否 | 否 | ✓ 1189ms | ✓ 1240ms | http |
| 38.145.203.96:8443 | ✓ 285ms | ✓ 836ms | ✓ 243ms | ✓ 852ms | ✓ 699ms | http |
| 38.145.208.244:8448 | ✓ 308ms | ✓ 919ms | ✓ 228ms | ✓ 821ms | ✓ 646ms | http |
| 38.145.208.247:8443 | ✓ 306ms | ✓ 889ms | ✓ 234ms | ✓ 841ms | ✓ 654ms | http |
| 38.145.218.210:8443 | ✓ 308ms | ✓ 831ms | ✓ 236ms | ✓ 873ms | ✓ 670ms | http |
| 38.145.203.86:8443 | ✓ 277ms | ✓ 1365ms | ✓ 243ms | ✓ 861ms | ✓ 659ms | http |
| 38.145.220.93:8443 | ✓ 292ms | ✓ 791ms | ✓ 241ms | ✓ 866ms | ✓ 878ms | http |
| 38.145.218.229:8450 | ✓ 245ms | ✓ 1372ms | ✓ 237ms | ✓ 875ms | ✓ 647ms | http |
| 38.34.179.50:8452 | ✓ 256ms | ✓ 857ms | ✓ 241ms | ✓ 875ms | ✓ 664ms | http |

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

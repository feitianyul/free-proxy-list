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

最后更新：2026-03-12 16:57:18 UTC（2026-03-13 00:57:18 UTC+8）

**代理总数：79**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 103.84.95.54:7890 | ✓ 860ms | 否 | ✓ 728ms | ✓ 790ms | ✓ 959ms | http |
| 205.209.118.30:3138 | ✓ 574ms | 否 | ✓ 1467ms | ✓ 1394ms | ✓ 1170ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1230ms | ✓ 1197ms | ✓ 997ms | http |
| 171.251.172.78:5106 | 否 | 否 | ✓ 1716ms | ✓ 1496ms | ✓ 1317ms | http |
| 217.76.245.80:999 | ✓ 838ms | ✓ 1613ms | ✓ 1310ms | 否 | ✓ 1416ms | http |
| 45.136.130.223:8443 | ✓ 915ms | 否 | ✓ 954ms | ✓ 1050ms | ✓ 1769ms | http |
| 101.47.73.135:3128 | ✓ 982ms | 否 | ✓ 1326ms | ✓ 1812ms | ✓ 1009ms | http |
| 103.139.138.194:3128 | ✓ 1752ms | 否 | ✓ 1582ms | ✓ 1311ms | ✓ 1146ms | http |
| 61.0.226.241:3128 | ✓ 1705ms | 否 | ✓ 1007ms | 否 | ✓ 1549ms | http |
| 190.9.109.198:999 | ✓ 1954ms | 否 | 否 | ✓ 1523ms | ✓ 1233ms | http |
| 8.219.97.248:80 | ✓ 881ms | 否 | ✓ 1228ms | 否 | ✓ 1503ms | http |
| 14.225.212.37:7890 | ✓ 1259ms | ✓ 1322ms | ✓ 855ms | ✓ 1720ms | ✓ 791ms | http |
| 107.173.52.58:7890 | ✓ 1493ms | 否 | ✓ 1172ms | 否 | ✓ 1269ms | http |
| 171.251.172.78:5104 | ✓ 1536ms | 否 | 否 | ✓ 1713ms | ✓ 1293ms | http |
| 45.136.131.47:8443 | ✓ 85ms | ✓ 1080ms | ✓ 823ms | ✓ 1281ms | ✓ 577ms | http |
| 117.159.239.51:22222 | ✓ 763ms | ✓ 1130ms | ✓ 911ms | ✓ 1042ms | ✓ 916ms | http |
| 152.42.213.210:8080 | ✓ 836ms | 否 | ✓ 1039ms | ✓ 1028ms | ✓ 837ms | http |
| 120.92.212.16:7890 | ✓ 895ms | ✓ 1180ms | ✓ 916ms | ✓ 1388ms | 否 | http |
| 120.92.212.16:8890 | ✓ 903ms | 否 | 否 | ✓ 1146ms | ✓ 903ms | http |
| 81.70.169.194:80 | ✓ 1023ms | 否 | ✓ 1188ms | ✓ 1247ms | 否 | http |
| 101.43.255.96:80 | 否 | ✓ 1406ms | ✓ 1649ms | 否 | ✓ 932ms | http |
| 59.46.216.131:30001 | ✓ 914ms | 否 | 否 | ✓ 1679ms | ✓ 1926ms | http |
| 120.238.159.189:22222 | ✓ 1370ms | ✓ 1131ms | 否 | ✓ 1082ms | ✓ 821ms | http |
| 113.59.32.161:22222 | ✓ 1586ms | ✓ 1308ms | ✓ 1123ms | 否 | ✓ 952ms | http |
| 46.183.25.8:443 | 否 | 否 | ✓ 1732ms | ✓ 989ms | ✓ 811ms | http |
| 120.198.141.75:22222 | 否 | ✓ 1215ms | ✓ 1015ms | ✓ 1194ms | ✓ 937ms | http |
| 113.59.32.142:22222 | ✓ 1020ms | ✓ 1278ms | ✓ 1100ms | ✓ 1164ms | 否 | http |
| 45.136.198.40:3128 | ✓ 860ms | 否 | ✓ 830ms | ✓ 1713ms | ✓ 1363ms | http |
| 150.107.140.238:3128 | ✓ 1785ms | 否 | ✓ 1921ms | ✓ 1090ms | ✓ 921ms | http |
| 114.231.73.92:1080 | 否 | ✓ 1223ms | 否 | ✓ 1489ms | ✓ 1729ms | http |
| 168.235.110.63:3128 | ✓ 799ms | ✓ 1998ms | ✓ 1927ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1313ms | ✓ 1489ms | 否 | 否 | ✓ 921ms | http |
| 116.80.47.62:3172 | ✓ 1993ms | 否 | 否 | ✓ 1930ms | ✓ 1629ms | http |
| 101.32.244.83:8080 | ✓ 988ms | 否 | ✓ 890ms | ✓ 1219ms | ✓ 1201ms | http |
| 121.43.196.213:8222 | ✓ 943ms | ✓ 1026ms | ✓ 836ms | ✓ 1066ms | ✓ 826ms | http |
| 121.43.196.210:8222 | ✓ 927ms | ✓ 1070ms | ✓ 810ms | ✓ 1052ms | ✓ 858ms | http |
| 114.55.226.123:10086 | ✓ 1677ms | 否 | ✓ 952ms | ✓ 1235ms | ✓ 994ms | http |
| 45.10.69.30:8888 | ✓ 1934ms | ✓ 916ms | ✓ 591ms | ✓ 670ms | ✓ 534ms | http |
| 116.80.95.238:7777 | 否 | 否 | ✓ 1479ms | ✓ 1779ms | ✓ 1697ms | http |
| 47.77.193.180:1080 | ✓ 813ms | 否 | ✓ 307ms | ✓ 1592ms | ✓ 491ms | http |
| 103.113.70.189:1081 | ✓ 1318ms | 否 | ✓ 967ms | ✓ 1324ms | ✓ 988ms | http |
| 45.136.130.175:8443 | ✓ 732ms | ✓ 605ms | ✓ 299ms | ✓ 667ms | ✓ 530ms | http |
| 120.240.35.173:22222 | ✓ 969ms | ✓ 1147ms | ✓ 933ms | ✓ 1140ms | ✓ 919ms | http |
| 120.240.29.51:22222 | 否 | ✓ 1179ms | ✓ 1017ms | ✓ 1092ms | ✓ 869ms | http |
| 117.159.239.44:22222 | ✓ 940ms | ✓ 1020ms | ✓ 798ms | ✓ 1015ms | ✓ 806ms | http |
| 183.249.5.110:22222 | ✓ 876ms | ✓ 813ms | ✓ 732ms | ✓ 975ms | ✓ 718ms | http |
| 45.236.129.64:3128 | 否 | ✓ 1916ms | ✓ 1096ms | 否 | ✓ 1648ms | http |
| 62.113.119.14:8080 | ✓ 1234ms | 否 | ✓ 1187ms | 否 | ✓ 1495ms | http |
| 180.127.149.244:1080 | ✓ 982ms | ✓ 1137ms | ✓ 858ms | 否 | ✓ 1211ms | http |
| 165.227.5.10:8888 | ✓ 307ms | ✓ 663ms | 否 | ✓ 1248ms | 否 | http |
| 35.225.22.61:80 | ✓ 586ms | 否 | ✓ 1023ms | ✓ 1395ms | 否 | http |
| 54.222.174.194:80 | 否 | 否 | ✓ 1690ms | ✓ 1910ms | ✓ 1727ms | http |
| 124.16.111.161:7890 | ✓ 760ms | ✓ 994ms | ✓ 896ms | ✓ 1015ms | ✓ 810ms | http |
| 39.104.201.40:7890 | 否 | ✓ 1484ms | ✓ 889ms | ✓ 1154ms | ✓ 922ms | http |
| 93.174.125.63:80 | ✓ 1707ms | 否 | ✓ 1953ms | 否 | ✓ 1807ms | http |
| 45.136.131.63:8443 | ✓ 701ms | ✓ 1038ms | ✓ 79ms | ✓ 828ms | ✓ 660ms | http |
| 183.249.5.214:22222 | ✓ 696ms | ✓ 830ms | ✓ 708ms | ✓ 903ms | ✓ 669ms | http |
| 120.55.163.237:10086 | ✓ 841ms | ✓ 990ms | ✓ 908ms | ✓ 1085ms | ✓ 840ms | http |
| 103.39.51.190:8080 | ✓ 1827ms | 否 | 否 | ✓ 1320ms | ✓ 1237ms | http |
| 202.129.206.239:3128 | ✓ 1696ms | 否 | ✓ 1879ms | ✓ 1950ms | ✓ 1594ms | http |
| 183.249.5.109:22222 | ✓ 780ms | ✓ 827ms | ✓ 742ms | ✓ 887ms | ✓ 776ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 866ms | ✓ 1755ms | ✓ 1361ms | http |
| 91.107.141.42:8081 | ✓ 1207ms | 否 | ✓ 1319ms | ✓ 1964ms | 否 | http |
| 45.136.130.188:8443 | 否 | 否 | ✓ 1574ms | ✓ 1146ms | ✓ 647ms | http |
| 45.136.130.191:8443 | 否 | 否 | ✓ 1567ms | ✓ 1152ms | ✓ 652ms | http |
| 183.249.5.213:22222 | 否 | ✓ 813ms | ✓ 699ms | ✓ 882ms | 否 | http |
| 183.249.5.111:22222 | ✓ 682ms | ✓ 886ms | ✓ 640ms | ✓ 920ms | ✓ 689ms | http |
| 120.240.35.160:22222 | 否 | ✓ 1243ms | ✓ 1005ms | ✓ 1124ms | ✓ 912ms | http |
| 24.144.86.173:1080 | 否 | 否 | ✓ 1450ms | ✓ 1302ms | ✓ 1670ms | http |
| 121.230.9.26:1080 | 否 | ✓ 1656ms | ✓ 1130ms | ✓ 1470ms | 否 | http |
| 113.59.32.160:22222 | 否 | ✓ 1252ms | ✓ 933ms | ✓ 1233ms | ✓ 893ms | http |
| 120.240.35.176:22222 | 否 | ✓ 1129ms | ✓ 816ms | ✓ 1093ms | ✓ 883ms | http |
| 85.208.108.43:10808 | ✓ 989ms | 否 | ✓ 1008ms | ✓ 1349ms | 否 | http |
| 85.208.108.43:2094 | ✓ 1369ms | 否 | ✓ 465ms | ✓ 1339ms | 否 | http |
| 120.240.35.177:22222 | ✓ 1157ms | 否 | ✓ 1070ms | 否 | ✓ 863ms | http |
| 181.78.44.63:999 | ✓ 998ms | ✓ 1413ms | ✓ 1652ms | ✓ 1778ms | ✓ 1239ms | http |
| 162.240.154.26:3128 | ✓ 1511ms | 否 | ✓ 1548ms | ✓ 1752ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1944ms | ✓ 1934ms | ✓ 1531ms | ✓ 1790ms | ✓ 1641ms | http |
| 113.59.32.162:22222 | ✓ 1035ms | ✓ 1247ms | ✓ 967ms | ✓ 1113ms | ✓ 927ms | http |

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

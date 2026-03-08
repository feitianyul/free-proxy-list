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

最后更新：2026-03-08 22:21:31 UTC（2026-03-09 06:21:31 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 81 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 152.42.213.210:8080 | ✓ 1445ms | 否 | ✓ 1098ms | 否 | ✓ 1093ms | http |
| 162.248.165.72:1080 | ✓ 906ms | 否 | 否 | ✓ 1914ms | ✓ 1952ms | http |
| 205.209.118.30:3138 | ✓ 443ms | 否 | ✓ 930ms | ✓ 1279ms | ✓ 975ms | http |
| 217.76.245.80:999 | ✓ 770ms | ✓ 1736ms | ✓ 1146ms | ✓ 1562ms | ✓ 1406ms | http |
| 45.116.14.87:8080 | ✓ 1370ms | ✓ 1471ms | ✓ 929ms | ✓ 1061ms | ✓ 768ms | http |
| 202.155.12.161:443 | 否 | ✓ 1404ms | 否 | ✓ 1312ms | ✓ 1521ms | http |
| 67.169.98.211:443 | ✓ 1201ms | 否 | ✓ 1982ms | ✓ 1965ms | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1087ms | 否 | ✓ 1953ms | ✓ 1164ms | http |
| 35.225.22.61:80 | ✓ 1485ms | ✓ 1289ms | ✓ 375ms | ✓ 1130ms | ✓ 874ms | http |
| 117.159.239.49:22222 | ✓ 932ms | ✓ 1066ms | ✓ 920ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1028ms | 否 | ✓ 1412ms | ✓ 1305ms | 否 | http |
| 45.77.246.231:80 | ✓ 1352ms | 否 | ✓ 1818ms | ✓ 1214ms | ✓ 1062ms | http |
| 62.113.119.14:8080 | ✓ 789ms | 否 | ✓ 801ms | ✓ 1783ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1910ms | ✓ 1921ms | 否 | ✓ 1767ms | ✓ 1598ms | http |
| 165.227.5.10:8888 | ✓ 1116ms | 否 | 否 | ✓ 1979ms | ✓ 1302ms | http |
| 45.186.6.104:3128 | ✓ 1796ms | ✓ 1934ms | ✓ 1694ms | 否 | 否 | http |
| 168.235.110.63:3128 | ✓ 1511ms | 否 | ✓ 1225ms | ✓ 1156ms | ✓ 900ms | http |
| 101.43.255.96:80 | ✓ 974ms | 否 | ✓ 1322ms | 否 | ✓ 1297ms | http |
| 81.70.169.194:80 | ✓ 1368ms | 否 | ✓ 1924ms | ✓ 1237ms | ✓ 1078ms | http |
| 190.9.109.207:999 | ✓ 864ms | ✓ 1494ms | ✓ 1093ms | ✓ 1397ms | ✓ 1334ms | http |
| 190.9.109.198:999 | ✓ 1075ms | ✓ 1539ms | ✓ 1223ms | ✓ 1330ms | ✓ 1070ms | http |
| 117.159.239.51:22222 | ✓ 885ms | ✓ 1064ms | ✓ 861ms | ✓ 1095ms | ✓ 849ms | http |
| 120.240.35.178:22222 | ✓ 963ms | ✓ 1328ms | ✓ 1043ms | ✓ 1142ms | ✓ 886ms | http |
| 46.249.103.192:443 | ✓ 675ms | 否 | ✓ 1588ms | ✓ 1871ms | ✓ 1921ms | http |
| 103.215.36.88:18925 | 否 | ✓ 1992ms | ✓ 1295ms | 否 | ✓ 1079ms | http |
| 180.76.115.231:3128 | ✓ 1386ms | ✓ 1618ms | ✓ 1085ms | ✓ 1404ms | ✓ 1146ms | http |
| 39.104.201.40:7890 | ✓ 937ms | ✓ 1204ms | ✓ 945ms | ✓ 1191ms | ✓ 943ms | http |
| 47.77.193.180:1080 | ✓ 397ms | ✓ 854ms | ✓ 439ms | ✓ 718ms | ✓ 549ms | http |
| 121.230.8.158:1080 | ✓ 1440ms | 否 | ✓ 1216ms | ✓ 1314ms | ✓ 1160ms | http |
| 117.159.239.48:22222 | 否 | ✓ 1100ms | 否 | ✓ 1133ms | ✓ 914ms | http |
| 120.92.212.16:8890 | ✓ 1350ms | ✓ 1520ms | ✓ 965ms | ✓ 1740ms | ✓ 1191ms | http |
| 147.45.251.242:8888 | ✓ 1909ms | 否 | ✓ 1940ms | 否 | ✓ 1979ms | http |
| 159.223.42.219:3128 | ✓ 907ms | 否 | ✓ 1366ms | ✓ 1442ms | ✓ 1936ms | http |
| 91.233.223.147:3128 | ✓ 1116ms | 否 | ✓ 1084ms | 否 | ✓ 1731ms | http |
| 117.159.239.52:22222 | ✓ 848ms | ✓ 1042ms | ✓ 885ms | ✓ 1114ms | ✓ 838ms | http |
| 222.184.48.242:22222 | ✓ 904ms | 否 | ✓ 909ms | ✓ 1199ms | ✓ 947ms | http |
| 120.240.29.168:22222 | ✓ 930ms | ✓ 1263ms | ✓ 1027ms | ✓ 1157ms | 否 | http |
| 222.184.48.252:22222 | ✓ 957ms | ✓ 1131ms | ✓ 958ms | ✓ 1161ms | ✓ 948ms | http |
| 120.240.35.177:22222 | ✓ 993ms | ✓ 1305ms | ✓ 975ms | ✓ 1170ms | ✓ 908ms | http |
| 113.59.32.163:22222 | ✓ 1075ms | ✓ 1328ms | 否 | ✓ 1293ms | ✓ 1042ms | http |
| 143.189.3.198:8080 | 否 | 否 | ✓ 580ms | ✓ 1831ms | ✓ 677ms | http |
| 101.32.244.83:8080 | ✓ 997ms | 否 | ✓ 948ms | ✓ 1196ms | ✓ 1223ms | http |
| 121.43.196.210:8222 | ✓ 986ms | ✓ 1098ms | ✓ 851ms | ✓ 1182ms | ✓ 921ms | http |
| 121.43.196.213:8222 | ✓ 991ms | ✓ 1082ms | ✓ 862ms | ✓ 1162ms | 否 | http |
| 114.55.226.123:10086 | ✓ 1075ms | ✓ 1391ms | ✓ 1096ms | ✓ 1351ms | ✓ 1091ms | http |
| 88.80.150.82:8080 | ✓ 1776ms | ✓ 1700ms | 否 | 否 | ✓ 1794ms | https |
| 1.231.81.166:3128 | ✓ 1190ms | ✓ 1589ms | ✓ 1132ms | ✓ 1561ms | ✓ 1194ms | http |
| 103.82.23.118:5216 | ✓ 1859ms | 否 | ✓ 1602ms | ✓ 1461ms | ✓ 1452ms | http |
| 116.80.82.232:3172 | ✓ 1649ms | 否 | 否 | ✓ 1824ms | ✓ 1682ms | http |
| 178.236.245.59:3128 | ✓ 1175ms | 否 | ✓ 1555ms | ✓ 1964ms | ✓ 1875ms | http |
| 194.213.18.200:443 | ✓ 1892ms | ✓ 1979ms | 否 | 否 | ✓ 1621ms | http |
| 178.236.245.17:3128 | ✓ 1129ms | 否 | ✓ 1856ms | ✓ 1753ms | ✓ 1871ms | http |
| 222.184.48.241:22222 | ✓ 960ms | 否 | ✓ 863ms | ✓ 1150ms | ✓ 870ms | http |
| 138.124.93.82:1080 | ✓ 1163ms | 否 | ✓ 1780ms | 否 | ✓ 1912ms | http |
| 103.215.36.88:18686 | ✓ 1088ms | ✓ 1195ms | 否 | ✓ 1402ms | ✓ 1085ms | http |
| 103.183.10.169:3125 | ✓ 1827ms | 否 | 否 | ✓ 1437ms | ✓ 1381ms | http |
| 83.219.250.8:62920 | ✓ 1255ms | 否 | ✓ 1594ms | 否 | ✓ 1712ms | http |
| 45.22.209.157:8888 | ✓ 1751ms | 否 | ✓ 1244ms | 否 | ✓ 1381ms | http |
| 172.212.68.37:3128 | ✓ 1233ms | 否 | 否 | ✓ 1791ms | ✓ 1496ms | http |
| 51.79.207.21:8080 | ✓ 1262ms | 否 | ✓ 1960ms | ✓ 1994ms | ✓ 1198ms | http |
| 103.39.51.190:8080 | ✓ 1824ms | 否 | 否 | ✓ 1415ms | ✓ 1628ms | http |
| 14.56.107.244:3128 | ✓ 1507ms | ✓ 1118ms | ✓ 688ms | 否 | 否 | http |
| 107.172.125.217:3128 | ✓ 442ms | 否 | ✓ 715ms | ✓ 824ms | 否 | http |
| 103.215.36.88:18038 | ✓ 1754ms | ✓ 1188ms | ✓ 1011ms | ✓ 1316ms | ✓ 1035ms | http |
| 120.240.35.176:22222 | ✓ 950ms | 否 | ✓ 1043ms | ✓ 1175ms | 否 | http |
| 47.101.159.19:8899 | ✓ 1352ms | ✓ 1537ms | 否 | ✓ 1920ms | 否 | http |
| 112.78.187.186:8080 | 否 | 否 | ✓ 1849ms | ✓ 1432ms | ✓ 1409ms | http |
| 1.225.116.115:1080 | ✓ 1129ms | ✓ 948ms | ✓ 869ms | ✓ 904ms | ✓ 723ms | http |
| 45.136.198.40:3128 | ✓ 1587ms | ✓ 1719ms | 否 | 否 | ✓ 1883ms | http |
| 211.171.114.154:3128 | ✓ 1828ms | 否 | ✓ 1079ms | ✓ 1396ms | ✓ 1463ms | http |
| 103.183.10.203:3125 | 否 | 否 | ✓ 1303ms | ✓ 1541ms | ✓ 1833ms | http |
| 61.72.221.94:3128 | ✓ 804ms | ✓ 1579ms | 否 | ✓ 1039ms | 否 | http |
| 159.89.31.62:8080 | ✓ 553ms | 否 | ✓ 1997ms | ✓ 1963ms | ✓ 1734ms | http |
| 120.240.35.161:22222 | ✓ 959ms | ✓ 1448ms | ✓ 998ms | ✓ 1119ms | 否 | http |
| 183.249.5.213:22222 | ✓ 769ms | ✓ 968ms | 否 | 否 | ✓ 724ms | http |
| 157.0.142.246:10057 | ✓ 960ms | ✓ 1446ms | ✓ 965ms | ✓ 1297ms | ✓ 963ms | http |
| 103.84.95.54:7890 | ✓ 695ms | 否 | ✓ 1117ms | ✓ 1089ms | 否 | http |
| 103.215.36.88:15088 | ✓ 1061ms | 否 | ✓ 1141ms | ✓ 1548ms | ✓ 1177ms | http |
| 120.240.35.160:22222 | ✓ 930ms | ✓ 1180ms | ✓ 983ms | ✓ 1186ms | ✓ 942ms | http |
| 34.101.184.164:3128 | ✓ 1620ms | 否 | ✓ 861ms | ✓ 1219ms | ✓ 960ms | http |
| 103.215.36.88:16909 | ✓ 1307ms | ✓ 1337ms | 否 | ✓ 1360ms | 否 | http |
| 103.215.36.88:14435 | ✓ 1283ms | ✓ 1165ms | 否 | ✓ 1184ms | 否 | http |

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

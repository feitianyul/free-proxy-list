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

最后更新：2026-03-10 15:50:14 UTC（2026-03-10 23:50:14 UTC+8）

**代理总数：72**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.47:8443 | ✓ 959ms | 否 | ✓ 257ms | ✓ 956ms | ✓ 943ms | http |
| 205.209.118.30:3138 | ✓ 460ms | 否 | ✓ 1272ms | ✓ 1341ms | ✓ 1030ms | http |
| 154.3.236.202:3128 | ✓ 561ms | 否 | ✓ 1167ms | ✓ 1749ms | ✓ 1152ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1541ms | ✓ 884ms | ✓ 963ms | ✓ 737ms | http |
| 202.155.12.161:443 | ✓ 1684ms | 否 | 否 | ✓ 1955ms | ✓ 1294ms | http |
| 35.225.22.61:80 | ✓ 551ms | 否 | ✓ 934ms | ✓ 1373ms | ✓ 1144ms | http |
| 45.140.147.82:1081 | ✓ 1576ms | ✓ 1772ms | ✓ 1585ms | ✓ 1890ms | ✓ 1273ms | http |
| 217.76.245.80:999 | ✓ 1429ms | 否 | ✓ 1334ms | ✓ 1541ms | ✓ 1285ms | http |
| 115.231.181.40:8128 | ✓ 947ms | 否 | ✓ 1036ms | 否 | ✓ 1573ms | http |
| 120.92.212.16:7890 | ✓ 954ms | ✓ 1338ms | ✓ 1172ms | ✓ 1366ms | 否 | http |
| 47.77.193.180:1080 | ✓ 1795ms | ✓ 844ms | ✓ 489ms | ✓ 803ms | ✓ 1139ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1265ms | 否 | ✓ 1344ms | ✓ 1090ms | http |
| 81.70.169.194:80 | ✓ 1429ms | 否 | ✓ 996ms | 否 | ✓ 1659ms | http |
| 101.43.255.96:80 | ✓ 895ms | ✓ 1401ms | ✓ 1694ms | 否 | ✓ 1900ms | http |
| 101.47.73.135:3128 | ✓ 1474ms | 否 | ✓ 1499ms | ✓ 1192ms | ✓ 1011ms | http |
| 192.227.137.65:5050 | ✓ 147ms | 否 | ✓ 146ms | ✓ 846ms | ✓ 678ms | http |
| 158.69.185.37:3129 | ✓ 701ms | ✓ 1793ms | ✓ 1039ms | ✓ 1776ms | ✓ 1003ms | http |
| 190.212.131.238:3128 | ✓ 616ms | 否 | ✓ 890ms | ✓ 1591ms | 否 | http |
| 162.248.165.72:1080 | ✓ 1548ms | ✓ 1895ms | ✓ 635ms | 否 | ✓ 1952ms | http |
| 45.136.130.223:8443 | ✓ 1797ms | 否 | ✓ 249ms | ✓ 679ms | ✓ 505ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1576ms | ✓ 1777ms | ✓ 571ms | http |
| 14.225.212.37:7890 | ✓ 1954ms | 否 | ✓ 1210ms | 否 | ✓ 1225ms | http |
| 185.191.236.162:3128 | ✓ 1071ms | 否 | 否 | ✓ 1760ms | ✓ 1234ms | http |
| 183.249.5.214:22222 | 否 | ✓ 877ms | ✓ 1164ms | ✓ 1082ms | ✓ 726ms | http |
| 106.14.203.63:3333 | ✓ 922ms | 否 | ✓ 957ms | ✓ 1200ms | ✓ 856ms | http |
| 114.55.226.123:10086 | ✓ 1028ms | ✓ 1693ms | ✓ 1113ms | ✓ 1335ms | ✓ 1042ms | http |
| 46.183.25.8:443 | ✓ 603ms | 否 | ✓ 180ms | ✓ 717ms | ✓ 710ms | http |
| 45.140.147.82:1082 | ✓ 592ms | ✓ 1830ms | ✓ 931ms | 否 | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1493ms | ✓ 1454ms | ✓ 967ms | http |
| 45.136.198.40:3128 | ✓ 825ms | 否 | ✓ 1352ms | ✓ 1672ms | ✓ 1318ms | http |
| 162.240.154.26:3128 | ✓ 1120ms | 否 | ✓ 695ms | ✓ 1841ms | ✓ 1219ms | http |
| 113.59.32.163:22222 | ✓ 1057ms | ✓ 1297ms | ✓ 1009ms | ✓ 1290ms | ✓ 1072ms | http |
| 120.240.29.173:22222 | ✓ 1745ms | ✓ 1300ms | ✓ 1011ms | ✓ 1143ms | ✓ 1673ms | http |
| 193.168.173.136:443 | ✓ 1893ms | ✓ 1965ms | ✓ 1176ms | ✓ 1987ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1000ms | 否 | ✓ 942ms | ✓ 1275ms | ✓ 1220ms | http |
| 121.43.196.210:8222 | ✓ 989ms | ✓ 1065ms | ✓ 891ms | ✓ 1153ms | ✓ 860ms | http |
| 121.43.196.213:8222 | ✓ 952ms | ✓ 1082ms | ✓ 919ms | ✓ 1107ms | ✓ 878ms | http |
| 114.231.72.27:1080 | ✓ 1017ms | ✓ 1157ms | ✓ 1157ms | 否 | ✓ 1759ms | http |
| 117.159.239.58:22222 | ✓ 1228ms | ✓ 1144ms | ✓ 834ms | 否 | 否 | http |
| 39.104.201.40:7890 | 否 | ✓ 1318ms | ✓ 969ms | 否 | ✓ 946ms | http |
| 116.80.82.216:3172 | 否 | 否 | ✓ 1551ms | ✓ 1855ms | ✓ 1651ms | http |
| 116.80.82.224:3172 | ✓ 1507ms | 否 | ✓ 1926ms | ✓ 1889ms | 否 | http |
| 152.42.213.210:8080 | ✓ 722ms | 否 | ✓ 1217ms | ✓ 1041ms | ✓ 820ms | http |
| 103.82.23.118:5178 | ✓ 1961ms | 否 | 否 | ✓ 1824ms | ✓ 1557ms | http |
| 194.213.18.200:443 | ✓ 741ms | 否 | 否 | ✓ 1224ms | ✓ 1310ms | http |
| 178.236.245.17:3128 | ✓ 768ms | 否 | ✓ 864ms | 否 | ✓ 1541ms | http |
| 211.171.114.154:3128 | ✓ 1717ms | 否 | ✓ 1456ms | ✓ 1278ms | ✓ 1025ms | http |
| 217.77.102.18:3128 | ✓ 1457ms | 否 | ✓ 1613ms | 否 | ✓ 1826ms | http |
| 183.249.5.111:22222 | 否 | ✓ 891ms | ✓ 671ms | 否 | ✓ 998ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1217ms | ✓ 1532ms | ✓ 1435ms | http |
| 159.223.42.219:3128 | ✓ 848ms | 否 | ✓ 1441ms | ✓ 1069ms | ✓ 948ms | http |
| 185.41.152.110:3128 | ✓ 1367ms | ✓ 1986ms | ✓ 1819ms | 否 | ✓ 1952ms | http |
| 95.3.9.78:3128 | 否 | ✓ 1964ms | ✓ 1224ms | 否 | ✓ 1426ms | http |
| 34.101.184.164:3128 | ✓ 1720ms | 否 | ✓ 1420ms | ✓ 1495ms | ✓ 1105ms | http |
| 59.46.216.131:30001 | ✓ 1098ms | ✓ 1286ms | ✓ 1326ms | 否 | ✓ 1030ms | http |
| 111.79.111.126:3128 | ✓ 1718ms | 否 | 否 | ✓ 1395ms | ✓ 1431ms | http |
| 103.82.36.237:8888 | ✓ 1823ms | 否 | ✓ 1574ms | ✓ 1898ms | ✓ 1494ms | http |
| 152.70.98.46:8888 | ✓ 1691ms | 否 | 否 | ✓ 1166ms | ✓ 1055ms | http |
| 61.52.131.172:8443 | ✓ 862ms | ✓ 1149ms | ✓ 893ms | ✓ 1133ms | ✓ 912ms | http |
| 91.107.141.42:8081 | ✓ 680ms | 否 | 否 | ✓ 1956ms | ✓ 1630ms | http |
| 120.232.242.119:22222 | ✓ 868ms | ✓ 1161ms | ✓ 1189ms | ✓ 1597ms | ✓ 1725ms | http |
| 192.227.137.63:5050 | ✓ 525ms | 否 | ✓ 1777ms | ✓ 1290ms | ✓ 892ms | http |
| 121.138.61.193:8143 | 否 | ✓ 1469ms | 否 | ✓ 1030ms | ✓ 822ms | http |
| 120.238.159.228:22222 | ✓ 991ms | ✓ 1317ms | ✓ 1013ms | ✓ 1261ms | ✓ 914ms | http |
| 120.238.159.230:22222 | ✓ 998ms | ✓ 1230ms | ✓ 1018ms | ✓ 1114ms | ✓ 1565ms | http |
| 139.162.46.62:3128 | ✓ 1876ms | ✓ 1937ms | ✓ 1204ms | ✓ 1674ms | ✓ 920ms | http |
| 103.3.246.71:3128 | ✓ 1801ms | 否 | ✓ 1054ms | ✓ 1450ms | ✓ 1531ms | http |
| 210.223.44.230:3128 | ✓ 1629ms | 否 | ✓ 1317ms | ✓ 1870ms | 否 | http |
| 120.198.141.84:22222 | 否 | ✓ 1550ms | ✓ 1049ms | ✓ 1153ms | 否 | http |
| 113.59.32.162:22222 | ✓ 1101ms | ✓ 1293ms | ✓ 1063ms | 否 | ✓ 1037ms | http |
| 120.240.35.177:22222 | ✓ 905ms | ✓ 1221ms | ✓ 1034ms | ✓ 1146ms | ✓ 932ms | http |
| 120.240.35.176:22222 | ✓ 1131ms | ✓ 1197ms | ✓ 978ms | ✓ 1148ms | ✓ 1076ms | http |

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

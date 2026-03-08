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

最后更新：2026-03-08 19:35:26 UTC（2026-03-09 03:35:26 UTC+8）

**代理总数：68**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 1057ms | 否 | ✓ 1061ms | ✓ 1382ms | ✓ 1040ms | http |
| 168.235.110.63:3128 | ✓ 748ms | 否 | ✓ 979ms | ✓ 1664ms | ✓ 1040ms | http |
| 121.128.121.54:3128 | ✓ 1604ms | ✓ 1749ms | 否 | ✓ 1112ms | ✓ 754ms | http |
| 120.92.212.16:7890 | ✓ 914ms | 否 | ✓ 882ms | ✓ 1823ms | 否 | http |
| 217.76.245.80:999 | ✓ 1039ms | ✓ 1656ms | ✓ 1262ms | ✓ 1467ms | ✓ 1554ms | http |
| 14.56.107.244:3128 | ✓ 1667ms | ✓ 1912ms | ✓ 541ms | ✓ 978ms | ✓ 739ms | http |
| 152.42.213.210:8080 | ✓ 723ms | 否 | ✓ 1470ms | ✓ 1140ms | ✓ 919ms | http |
| 202.155.12.161:443 | ✓ 1733ms | 否 | 否 | ✓ 1801ms | ✓ 1535ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 738ms | ✓ 1278ms | ✓ 502ms | http |
| 117.159.239.49:22222 | ✓ 907ms | ✓ 1046ms | ✓ 884ms | ✓ 1026ms | ✓ 828ms | http |
| 194.213.18.200:443 | ✓ 1633ms | ✓ 1980ms | 否 | 否 | ✓ 1755ms | http |
| 210.223.44.230:3128 | ✓ 655ms | ✓ 976ms | ✓ 720ms | ✓ 967ms | ✓ 776ms | http |
| 101.43.255.96:80 | ✓ 1021ms | ✓ 1301ms | ✓ 970ms | ✓ 1284ms | ✓ 997ms | http |
| 39.104.201.40:7890 | ✓ 878ms | ✓ 1146ms | ✓ 959ms | ✓ 1155ms | ✓ 1693ms | http |
| 81.70.169.194:80 | ✓ 931ms | ✓ 1180ms | ✓ 901ms | ✓ 1185ms | ✓ 1839ms | http |
| 190.9.109.207:999 | ✓ 810ms | ✓ 1692ms | ✓ 1272ms | 否 | ✓ 1454ms | http |
| 113.59.32.161:22222 | ✓ 992ms | ✓ 1304ms | ✓ 988ms | 否 | ✓ 908ms | http |
| 120.240.35.160:22222 | ✓ 859ms | ✓ 1188ms | ✓ 886ms | ✓ 1105ms | ✓ 925ms | http |
| 115.231.181.40:8128 | ✓ 896ms | ✓ 1011ms | ✓ 838ms | ✓ 1196ms | ✓ 963ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1108ms | 否 | ✓ 1360ms | ✓ 1120ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1152ms | 否 | ✓ 1141ms | ✓ 951ms | http |
| 103.215.36.88:10855 | 否 | ✓ 1159ms | ✓ 970ms | 否 | ✓ 922ms | http |
| 222.184.48.241:22222 | ✓ 1205ms | ✓ 1208ms | ✓ 828ms | ✓ 1093ms | ✓ 948ms | http |
| 61.72.221.94:3128 | ✓ 1581ms | ✓ 1607ms | ✓ 1217ms | ✓ 913ms | ✓ 797ms | http |
| 125.128.12.14:3128 | ✓ 1582ms | ✓ 1175ms | ✓ 1859ms | ✓ 1220ms | 否 | http |
| 211.171.114.154:3128 | 否 | 否 | ✓ 1779ms | ✓ 1525ms | ✓ 974ms | http |
| 47.77.193.180:1080 | ✓ 371ms | ✓ 1293ms | ✓ 197ms | ✓ 668ms | ✓ 525ms | http |
| 62.113.119.14:8080 | ✓ 977ms | ✓ 1702ms | ✓ 923ms | ✓ 1692ms | 否 | http |
| 152.42.213.210:80 | ✓ 840ms | 否 | ✓ 1156ms | ✓ 1372ms | ✓ 1488ms | http |
| 45.140.147.82:1081 | ✓ 729ms | 否 | ✓ 1696ms | 否 | ✓ 1174ms | http |
| 1.231.81.166:3128 | ✓ 1713ms | ✓ 792ms | ✓ 677ms | ✓ 838ms | ✓ 677ms | http |
| 35.225.22.61:80 | ✓ 678ms | 否 | ✓ 1054ms | ✓ 1186ms | ✓ 913ms | http |
| 150.249.255.91:3128 | ✓ 1628ms | ✓ 1821ms | ✓ 1775ms | 否 | ✓ 781ms | http |
| 183.249.5.111:22222 | ✓ 655ms | ✓ 1076ms | ✓ 674ms | ✓ 1105ms | ✓ 680ms | http |
| 117.159.239.44:22222 | ✓ 836ms | 否 | ✓ 829ms | ✓ 1051ms | ✓ 827ms | http |
| 117.159.239.51:22222 | ✓ 796ms | ✓ 981ms | ✓ 752ms | ✓ 1005ms | ✓ 835ms | http |
| 120.240.35.176:22222 | ✓ 898ms | ✓ 1193ms | 否 | ✓ 1075ms | ✓ 836ms | http |
| 222.184.48.242:22222 | ✓ 905ms | 否 | ✓ 1246ms | ✓ 1170ms | ✓ 899ms | http |
| 113.59.32.142:22222 | ✓ 986ms | ✓ 1246ms | 否 | ✓ 1115ms | ✓ 927ms | http |
| 121.230.8.208:1080 | ✓ 1050ms | ✓ 1478ms | ✓ 1128ms | ✓ 1400ms | ✓ 1266ms | http |
| 178.236.245.17:3128 | ✓ 1266ms | ✓ 1747ms | ✓ 1942ms | 否 | ✓ 1739ms | http |
| 178.236.245.59:3128 | ✓ 1266ms | ✓ 1886ms | ✓ 1803ms | 否 | ✓ 1739ms | http |
| 120.240.35.177:22222 | ✓ 903ms | ✓ 1168ms | ✓ 967ms | 否 | 否 | http |
| 120.240.35.178:22222 | ✓ 1035ms | ✓ 1214ms | ✓ 961ms | ✓ 1149ms | 否 | http |
| 103.39.51.207:8080 | ✓ 1671ms | 否 | 否 | ✓ 1298ms | ✓ 1780ms | http |
| 101.32.244.83:8080 | ✓ 1361ms | 否 | ✓ 883ms | ✓ 1101ms | ✓ 990ms | http |
| 121.43.196.210:8222 | ✓ 912ms | ✓ 1025ms | ✓ 815ms | ✓ 1083ms | ✓ 858ms | http |
| 121.43.196.213:8222 | ✓ 996ms | ✓ 1001ms | ✓ 814ms | ✓ 1099ms | ✓ 828ms | http |
| 114.55.226.123:10086 | ✓ 1042ms | ✓ 1373ms | ✓ 1032ms | ✓ 1234ms | ✓ 1006ms | http |
| 8.219.97.248:80 | ✓ 1739ms | 否 | ✓ 1339ms | 否 | ✓ 1489ms | http |
| 125.128.12.144:3128 | ✓ 1662ms | 否 | ✓ 1180ms | ✓ 1488ms | ✓ 897ms | http |
| 45.22.209.157:8888 | ✓ 1301ms | 否 | ✓ 1458ms | ✓ 1359ms | ✓ 1347ms | http |
| 103.84.95.54:7890 | ✓ 828ms | 否 | ✓ 613ms | 否 | ✓ 621ms | http |
| 103.39.51.190:8080 | ✓ 1752ms | 否 | 否 | ✓ 1292ms | ✓ 1661ms | http |
| 45.186.6.104:3128 | ✓ 1966ms | ✓ 1630ms | ✓ 1741ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 711ms | ✓ 1461ms | ✓ 1334ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 910ms | ✓ 1932ms | ✓ 813ms | 否 | ✓ 1716ms | http |
| 61.72.110.54:3128 | ✓ 1512ms | ✓ 1620ms | ✓ 1067ms | ✓ 972ms | ✓ 797ms | http |
| 103.215.36.88:18169 | ✓ 1014ms | ✓ 1243ms | ✓ 998ms | 否 | ✓ 1049ms | http |
| 120.240.35.175:22222 | ✓ 855ms | ✓ 1185ms | ✓ 972ms | 否 | ✓ 922ms | http |
| 103.215.36.88:15412 | ✓ 1025ms | ✓ 1205ms | ✓ 1325ms | 否 | 否 | http |
| 120.240.35.173:22222 | ✓ 842ms | ✓ 1181ms | ✓ 1070ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 899ms | ✓ 1384ms | ✓ 1021ms | 否 | 否 | http |
| 120.198.141.75:22222 | ✓ 1157ms | ✓ 1532ms | ✓ 944ms | 否 | 否 | http |
| 113.59.32.163:22222 | ✓ 959ms | ✓ 1239ms | ✓ 1069ms | ✓ 1246ms | ✓ 982ms | http |
| 120.240.35.161:22222 | 否 | ✓ 1225ms | ✓ 967ms | ✓ 1090ms | ✓ 859ms | http |
| 172.212.68.37:3128 | ✓ 1044ms | ✓ 1953ms | ✓ 1369ms | 否 | ✓ 1613ms | http |
| 88.80.150.82:8080 | ✓ 1012ms | 否 | ✓ 1272ms | ✓ 1916ms | ✓ 1517ms | https |

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

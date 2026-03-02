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

最后更新：2026-03-02 18:47:44 UTC（2026-03-03 02:47:44 UTC+8）

**代理总数：53**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | 否 | 否 | ✓ 1585ms | ✓ 1178ms | ✓ 1084ms | http |
| 103.84.95.54:7890 | ✓ 1479ms | 否 | ✓ 662ms | ✓ 1295ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1205ms | 否 | ✓ 984ms | ✓ 1212ms | ✓ 954ms | http |
| 5.75.196.26:40000 | ✓ 573ms | ✓ 1710ms | ✓ 1998ms | 否 | ✓ 1071ms | http |
| 217.76.245.80:999 | ✓ 1402ms | 否 | 否 | ✓ 1390ms | ✓ 1192ms | http |
| 115.76.5.32:10010 | ✓ 1418ms | 否 | ✓ 1788ms | ✓ 1647ms | ✓ 1421ms | http |
| 205.209.118.30:3138 | 否 | ✓ 1154ms | ✓ 273ms | 否 | ✓ 1008ms | http |
| 223.16.170.103:80 | ✓ 1633ms | 否 | ✓ 1242ms | ✓ 1615ms | ✓ 1074ms | http |
| 61.72.110.54:3128 | ✓ 1249ms | 否 | 否 | ✓ 1068ms | ✓ 840ms | http |
| 101.43.255.96:80 | ✓ 1037ms | ✓ 1300ms | ✓ 877ms | ✓ 1237ms | ✓ 909ms | http |
| 81.70.169.194:80 | ✓ 944ms | ✓ 1786ms | ✓ 954ms | ✓ 1194ms | ✓ 1039ms | http |
| 45.129.141.143:3128 | ✓ 880ms | 否 | ✓ 1902ms | 否 | ✓ 1879ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1983ms | ✓ 1291ms | ✓ 958ms | http |
| 35.234.17.221:8080 | 否 | ✓ 1284ms | ✓ 1251ms | ✓ 992ms | 否 | http |
| 142.171.85.32:1080 | 否 | ✓ 1860ms | ✓ 1266ms | ✓ 1327ms | ✓ 1457ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1444ms | ✓ 1402ms | ✓ 1316ms | http |
| 91.99.99.83:9000 | ✓ 886ms | ✓ 1780ms | ✓ 1415ms | 否 | ✓ 1957ms | http |
| 138.124.53.25:7443 | ✓ 703ms | 否 | ✓ 1314ms | 否 | ✓ 1592ms | http |
| 120.92.212.16:7890 | ✓ 1943ms | ✓ 1259ms | 否 | 否 | ✓ 1182ms | http |
| 45.136.198.40:3128 | ✓ 1342ms | 否 | ✓ 1950ms | 否 | ✓ 1729ms | http |
| 115.76.5.32:10008 | ✓ 1326ms | 否 | ✓ 1871ms | 否 | ✓ 1342ms | http |
| 115.76.5.32:10009 | ✓ 1896ms | 否 | ✓ 1377ms | ✓ 1660ms | 否 | http |
| 103.215.36.88:17853 | ✓ 1621ms | ✓ 1837ms | 否 | 否 | ✓ 1553ms | http |
| 101.32.244.83:8080 | ✓ 1298ms | ✓ 1485ms | ✓ 925ms | ✓ 1271ms | ✓ 1236ms | http |
| 121.43.196.213:8222 | ✓ 915ms | ✓ 1088ms | ✓ 842ms | ✓ 1096ms | ✓ 921ms | http |
| 121.43.196.210:8222 | ✓ 894ms | ✓ 1061ms | ✓ 873ms | ✓ 1147ms | ✓ 940ms | http |
| 114.55.226.123:10086 | ✓ 1085ms | ✓ 1379ms | ✓ 1032ms | ✓ 1223ms | ✓ 983ms | http |
| 222.184.48.248:22222 | 否 | ✓ 1408ms | ✓ 886ms | ✓ 1222ms | ✓ 949ms | http |
| 46.249.103.192:443 | ✓ 1051ms | 否 | ✓ 1784ms | 否 | ✓ 1951ms | http |
| 45.140.147.82:1081 | ✓ 1070ms | 否 | 否 | ✓ 1735ms | ✓ 1857ms | http |
| 116.6.106.33:3128 | ✓ 713ms | ✓ 1004ms | ✓ 801ms | ✓ 882ms | ✓ 710ms | http |
| 103.215.36.88:10864 | ✓ 1003ms | 否 | ✓ 1133ms | ✓ 1372ms | ✓ 1103ms | http |
| 160.238.65.8:3128 | ✓ 613ms | ✓ 1927ms | ✓ 1802ms | 否 | ✓ 1713ms | http |
| 157.230.38.173:3128 | 否 | 否 | ✓ 1146ms | ✓ 1046ms | ✓ 865ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1484ms | 否 | ✓ 1815ms | ✓ 1216ms | http |
| 103.39.51.190:8080 | ✓ 1683ms | 否 | 否 | ✓ 1564ms | ✓ 1463ms | http |
| 115.76.5.32:10005 | ✓ 1979ms | 否 | ✓ 1888ms | ✓ 1794ms | 否 | http |
| 47.77.180.205:1080 | 否 | ✓ 738ms | ✓ 112ms | ✓ 765ms | ✓ 600ms | http |
| 183.249.5.109:22222 | ✓ 739ms | ✓ 853ms | ✓ 755ms | ✓ 920ms | ✓ 707ms | http |
| 113.59.32.145:22222 | ✓ 1096ms | ✓ 1376ms | ✓ 1121ms | ✓ 1255ms | ✓ 1001ms | http |
| 113.59.32.161:22222 | ✓ 1087ms | ✓ 1345ms | ✓ 1081ms | ✓ 1247ms | ✓ 978ms | http |
| 120.198.141.79:22222 | ✓ 1072ms | ✓ 1615ms | ✓ 1178ms | ✓ 1355ms | ✓ 1088ms | http |
| 120.240.29.174:22222 | ✓ 970ms | ✓ 1235ms | ✓ 881ms | ✓ 1148ms | ✓ 939ms | http |
| 120.232.242.119:22222 | ✓ 1805ms | ✓ 1182ms | ✓ 895ms | ✓ 1131ms | ✓ 904ms | http |
| 120.240.35.178:22222 | ✓ 1061ms | ✓ 1526ms | ✓ 1010ms | ✓ 1355ms | ✓ 1073ms | http |
| 34.101.184.164:3128 | ✓ 899ms | 否 | ✓ 1308ms | ✓ 1603ms | ✓ 1280ms | http |
| 85.208.108.43:2094 | ✓ 520ms | 否 | ✓ 993ms | 否 | ✓ 855ms | http |
| 120.238.159.228:22222 | ✓ 858ms | ✓ 1247ms | ✓ 1031ms | ✓ 1209ms | ✓ 912ms | http |
| 117.159.239.51:22222 | 否 | ✓ 1013ms | ✓ 802ms | ✓ 1078ms | ✓ 843ms | http |
| 113.59.32.142:22222 | ✓ 1161ms | ✓ 1429ms | ✓ 1068ms | ✓ 1272ms | ✓ 961ms | http |
| 117.159.239.44:22222 | ✓ 1010ms | ✓ 1248ms | ✓ 1081ms | 否 | ✓ 860ms | http |
| 120.198.141.80:22222 | ✓ 1122ms | ✓ 1590ms | 否 | ✓ 1259ms | 否 | http |
| 120.240.35.160:22222 | ✓ 1120ms | ✓ 1593ms | ✓ 1005ms | 否 | ✓ 1110ms | http |

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
